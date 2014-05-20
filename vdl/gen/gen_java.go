package gen

import (
	"bytes"
	"fmt"
	"log"
	"path"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"

	"veyron2/val"
	"veyron2/vdl/compile"
)

// javaGenPkgPrefix is the path prefix to be added to generated VDL package paths.
var javaGenPkgPrefix string = "com"

// SetJavaGenPkgPrefix sets the prefix that will be added to generated VDL package paths.
func SetJavaGenPkgPrefix(prefix string) {
	javaGenPkgPrefix = prefix
}

// javaGenPkgPath returns the Java package path given the Go package path.
func javaGenPkgPath(goPkgPath string) string {
	return path.Join(javaGenPkgPrefix, goPkgPath)
}

// JavaFileInfo stores the name and contents of the generated Java file.
type JavaFileInfo struct {
	Name string
	Data []byte
}

// GenJavaFiles generates Java files for all VDL files in the provided package,
// returning the list of generated Java files as a slice.  Since Java requires
// that each public class/interface gets defined in a separate file, this method
// will return one generated file per struct.  (Interfaces actually generate
// two files because we create separate interfaces for clients and services.)
// In addition, since Java doesn't support global variables (i.e., variables
// defined outside of a class), all constants are moved into a special "Consts"
// class and stored in a separate file.  All client bindings are stored in a
// separate Client.java file. Finally, package documentation (if any) is stored
// in a "package-info.java" file.
//
// The current generator doesn't yet support the full set of VDL features.  In
// particular, we don't yet support error ids and types Complex64 and Complex128.
//
// TODO(spetrovic): Run Java formatters on the generated files.
func GenJavaFiles(pkg *compile.Package, env *compile.Env) (ret []JavaFileInfo) {
	// One file for package documentation (if any).
	if g := genJavaPackageFile(pkg, env); g != nil {
		ret = append(ret, *g)
	}
	// Single file for all constants' definitions.
	ret = append(ret, genJavaConstFile(pkg, env))
	for _, file := range pkg.Files {
		// Separate file for all (struct) typedefs.
		for _, tdef := range file.TypeDefs {
			if tdef.Type.Kind() == val.Struct {
				ret = append(ret, genJavaClassFile(tdef, env))
			}
		}
		// Separate file for all interface definitions.
		for _, iface := range file.Interfaces {
			ret = append(ret, genJavaInterfaceFile(iface, false, env)) // client interface
			ret = append(ret, genJavaInterfaceFile(iface, true, env))  // service interface
		}
	}
	// Single file for all client stub implementations.
	ret = append(ret, genJavaClientFile(pkg, env))
	return
}

// genPackageFileJava generates the Java package info file, iff any package
// comments were specified in the package's VDL files.
func genJavaPackageFile(pkg *compile.Package, env *compile.Env) *JavaFileInfo {
	for _, file := range pkg.Files {
		if file.PackageDef.Doc != "" {
			var buf bytes.Buffer
			err := javaPkgTmpl.Execute(&buf, file)
			if err != nil {
				log.Fatalf("vdl: couldn't execute package template: %v", err)
			}
			return &JavaFileInfo{
				Name: "package-info.java",
				Data: buf.Bytes(),
			}
		}
	}
	return nil
}

const javaPkgTmplStr = `
{{$file := .}}// This file was auto-generated by the veyron vdl tool.
// Source: {{$file.BaseName}}

{{javaDoc $file.PackageDef.Doc}}
package {{javaPath (javaGenPkgPath $file.Package.Path)}};{{$file.PackageDef.DocSuffix}}
`

// genConstFileJava generates the (single) Java file that contains constant
// definitions from all the VDL files.
func genJavaConstFile(pkg *compile.Package, env *compile.Env) JavaFileInfo {
	data := struct {
		Package *compile.Package
		Imports javaUserImports
		Env     *compile.Env
	}{
		Package: pkg,
		Imports: javaConstImports(pkg, env),
		Env:     env,
	}
	var buf bytes.Buffer
	err := javaConstTmpl.Execute(&buf, data)
	if err != nil {
		log.Fatal("vdl: couldn't execute const template: %v", err)
	}
	return JavaFileInfo{
		Name: "Consts.java",
		Data: buf.Bytes(),
	}
}

const javaConstTmplStr = `{{$data := .}}// This file was auto-generated by the veyron vdl tool.
// Source(s):{{range $file := $data.Package.Files}} {{$file.BaseName}}{{end}}
package {{javaPath (javaGenPkgPath $data.Package.Path)}};

public class Consts { {{range $file := $data.Package.Files}}
	/* The following constants originate in file: {{$file.BaseName}}. */{{range $const := $file.ConstDefs}}
	{{$const.Doc}}public static final {{javaType $const.Value.Type false $data.Imports $data.Env}} {{toConstCase $const.Name}} = {{javaConstVal $const.Value}};{{$const.DocSuffix}}{{end}}{{end}}
}
`

// genStructFile generates the Java class file for the provided struct.
func genJavaClassFile(tdef *compile.TypeDef, env *compile.Env) JavaFileInfo {
	data := struct {
		TypeDef *compile.TypeDef
		Imports javaUserImports
		Env     *compile.Env
	}{
		TypeDef: tdef,
		Imports: javaStructImports(tdef, env),
		Env:     env,
	}
	var buf bytes.Buffer
	err := javaClassTmpl.Execute(&buf, data)
	if err != nil {
		log.Fatalf("vdl: couldn't execute struct template: %v", err)
	}
	return JavaFileInfo{
		Name: tdef.Name + ".java",
		Data: buf.Bytes(),
	}
}

const javaClassTmplStr = `
{{$data := .}}{{$typeDef := $data.TypeDef}}// This file was auto-generated by the veyron vdl tool.
// Source: {{$typeDef.File.BaseName}}
package {{javaPath (javaGenPkgPath $typeDef.File.Package.Path)}};
{{range $imp := javaImports $data.Imports $typeDef.File.Package}}
import {{$imp}};{{end}}

{{javaDoc $typeDef.Doc}}public class {{javaQualifiedName $typeDef.Name $typeDef.File $data.Imports}} { {{range $idx, $field := javaStructFields $typeDef.Type}}
	{{index $typeDef.FieldDoc $idx}}public {{javaType $field.Type false $data.Imports $data.Env}} {{toCamelCase $field.Name}};{{index $typeDef.FieldDocSuffix $idx}}{{end}}
}
`

// genJavaInterfaceFile generates the Java interface file for the provided
// interface.
func genJavaInterfaceFile(iface *compile.Interface, isService bool, env *compile.Env) JavaFileInfo {
	data := struct {
		Interface *compile.Interface
		IsService bool
		Imports   javaUserImports
		Env       *compile.Env
	}{
		Interface: iface,
		IsService: isService,
		Imports:   javaInterfaceImports(iface, isService, env),
		Env:       env,
	}
	var buf bytes.Buffer
	err := javaIfaceTmpl.Execute(&buf, data)
	if err != nil {
		log.Fatalf("vdl: couldn't execute struct template: %v", err)
	}
	name := iface.Name
	if isService {
		name += "Service"
	}
	name += ".java"
	return JavaFileInfo{
		Name: name,
		Data: buf.Bytes(),
	}
}

const javaIfaceTmplStr = `{{$data := .}}{{$iface := $data.Interface}}// This file was auto-generated by the veyron vdl tool.
// Source: {{$iface.File.BaseName}}
package {{javaPath (javaGenPkgPath $iface.File.Package.Path)}};
{{range $imp := javaImports $data.Imports $iface.File.Package}}
import {{$imp}};{{end}}

{{javaDoc $iface.Doc}}public interface {{javaQualifiedName (javaIfaceName $iface $data.IsService) $iface.File $data.Imports}} {{if gt (len $iface.Embeds) 0}}extends {{end}}{{range $idx, $eIface := $iface.Embeds}}{{if gt $idx 0}}, {{end}}{{javaQualifiedName (javaIfaceName $eIface $data.IsService) $eIface.File $data.Imports}}{{end}} { {{range $method := $iface.Methods}}{{if javaShouldPackOutArgs $method $data.IsService $data.Imports $data.Env}}
	// {{$method.Name}}Out packages output arguments for method {{$method.Name}}.
	// TODO(spetrovic): should this be the same type for the service and the client?  If so, it needs
	// to be defined outside of the class and in its own file (ugh!!).
	public class {{$method.Name}}Out { {{range $arg := javaNonErrorOutArgs $method}}
		{{$arg.Doc}}public {{javaType $arg.Type false $data.Imports $data.Env}} {{toCamelCase $arg.Name}};{{$arg.DocSuffix}}{{end}}
	}{{end}}
	{{$method.Doc}}	public {{javaOutArgType $method $data.IsService false $data.Imports $data.Env}} {{toCamelCase $method.Name}}({{javaInArgs $method $data.IsService true $data.Imports $data.Env}}) throws {{javaName "com/veyron2/ipc/VeyronException" $data.Imports}};{{$method.DocSuffix}}{{end}}
}
`

// genJavaClientFile generates the Java file containing client bindings for
// all interfaces in the provided package.
func genJavaClientFile(pkg *compile.Package, env *compile.Env) JavaFileInfo {
	data := struct {
		Package *compile.Package
		Imports javaUserImports
		Env     *compile.Env
	}{
		Package: pkg,
		Imports: javaClientImports(pkg, env),
		Env:     env,
	}
	var buf bytes.Buffer
	err := javaClientTmpl.Execute(&buf, data)
	if err != nil {
		log.Fatal("vdl: couldn't execute client template: %v", err)
	}
	return JavaFileInfo{
		Name: "Client.java",
		Data: buf.Bytes(),
	}
}

const javaClientTmplStr = `{{$data := .}}{{$clientClassName := javaName "com/veyron2/ipc/Client" $data.Imports}}{{$stringClassName := javaName "java/lang/String" $data.Imports}}// This file was auto-generated by the veyron vdl tool.
// Source(s): {{range $file := $data.Package.Files}} {{$file.BaseName}}{{end}}
package {{javaPath (javaGenPkgPath $data.Package.Path)}};
{{range $imp := javaImports $data.Imports $data.Package}}
import {{$imp}};{{end}}

public class Client { {{range $file := $data.Package.Files}}
	/* Bind methods for interfaces from VDL file: {{$file.BaseName}}. */{{range $iface := $file.Interfaces}}{{$ifaceClassName := javaQualifiedName $iface.Name $iface.File $data.Imports}}
	public static {{$ifaceClassName}} Bind{{$ifaceClassName}}({{$stringClassName}} name, {{$clientClassName}}.BindOption... opts) {
		// TODO(spetrovic): check bind options.{{range $eIface := $iface.Embeds}}{{$eIfaceClassName := javaQualifiedName $eIface.Name $eIface.File $data.Imports}}
		final {{$eIfaceClassName}} {{javaClassToVarName $eIfaceClassName}} = {{javaName (printf "%s/Client" (javaGenPkgPath $eIface.File.Package.Path)) $data.Imports}}.Bind{{$eIface.Name}}(name, opts);{{end}}
		final {{$clientClassName}} client = {{javaName "com/veyron2/runtime/RuntimeFactory" $data.Imports}}.getRuntime().getClient();
		return new {{$ifaceClassName}}Stub(client, name{{range $eIface := $iface.Embeds}}{{$eIfaceClassName := javaQualifiedName $eIface.Name $eIface.File $data.Imports}}, {{javaClassToVarName $eIfaceClassName}}{{end}});
	}{{end}}{{end}}
	{{range $file := $data.Package.Files}}
	/* Client stubs for interfaces in file: {{$file.BaseName}}. */{{range $iface := $file.Interfaces}}{{$ifaceClassName := javaQualifiedName $iface.Name $iface.File $data.Imports}}
	private static class {{$ifaceClassName}}Stub implements {{$ifaceClassName}} {
		private final {{$clientClassName}} client;
		private final {{$stringClassName}} name;{{range $idx, $eIface := $iface.Embeds}}{{$eIfaceClassName := javaQualifiedName $eIface.Name $eIface.File $data.Imports}}
		private final {{$eIfaceClassName}} {{javaClassToVarName $eIfaceClassName}};{{end}}

		{{$ifaceClassName}}Stub({{$clientClassName}} client, {{$stringClassName}} name{{range $eIface := $iface.Embeds}}{{$eIfaceClassName := javaQualifiedName $eIface.Name $eIface.File $data.Imports}}, {{$eIfaceClassName}} {{javaClassToVarName $eIfaceClassName}}{{end}}) {
			this.client = client;
			this.name = name;{{range $eIface := $iface.Embeds}}{{$eIfaceClassName := javaQualifiedName $eIface.Name $eIface.File $data.Imports}}{{$varName := javaClassToVarName $eIfaceClassName}}
			this.{{$varName}} = {{$varName}};{{end}}
		}{{if gt (len $iface.Methods) 0}}
		// Methods from interface {{$iface.Name}}.{{end}}{{range $method := $iface.Methods}}{{$outArgType := javaOutArgType $method false false $data.Imports $data.Env}}
		@Override
		public {{$outArgType}} {{toCamelCase $method.Name}}({{javaInArgs $method false false $data.Imports $data.Env}}) throws {{javaName "com/veyron2/ipc/VeyronException" $data.Imports}} {
			// Prepare input arguments.{{$objectClassName := javaName "java/lang/Object" $data.Imports}}{{$typeTokenClassName := javaName "com/google/common/reflect/TypeToken" $data.Imports}}
			final {{$objectClassName}}[] inArgs = new {{$objectClassName}}[{{len $method.InArgs}}];{{if gt (len $method.InArgs) 0}}
			int inIdx = 0;{{range $arg := $method.InArgs}}
			inArgs[inIdx++] = {{$arg.Name}};{{end}}{{end}}

			// Start the call.
			final {{$clientClassName}}.Call call = this.client.startCall(this.name, "{{$method.Name}}", inArgs, opts);
			{{if javaIsStreamingMethod $method}}
			// TODO(spetrovic): implement streaming.
			return null;{{else}}{{$outArgs := javaOutArgs $method}}
			// Prepare output arguments and finish the call.
			final {{$typeTokenClassName}}<?>[] resultTypes = new {{$typeTokenClassName}}<?>[{{len $outArgs}}]; {{if gt (len $outArgs) 0}}
			int outIdx = 0;{{range $arg := $outArgs}}
			resultTypes[outIdx++] = new {{$typeTokenClassName}}<{{javaType $arg.Type true $data.Imports $data.Env}}>() {};{{end}}{{end}}{{if eq (len $outArgs) 0}}
			call.finish(resultTypes);
			return;{{else}}
			final Object[] results = call.finish(resultTypes);{{if eq (len $outArgs) 1}}
			return ({{$outArgType}})results[0];{{else}}

			// Pack the results.
			final {{$outArgType}} ret = new {{$outArgType}}();
			int resultIdx = 0;{{range $arg := $outArgs}}
			ret.{{toCamelCase $arg.Name}} = ({{javaType $arg.Type true $data.Imports $data.Env}})results[resultIdx++];{{end}}
			return ret;{{end}}{{end}}{{end}}
		}{{end}}{{range $eIface := javaAllEmbeddedIfaces $iface}}{{$eIfaceClassName := javaQualifiedName $eIface.Name $eIface.File $data.Imports}}{{if gt (len $eIface.Methods) 0}}
		// Methods from sub-interface {{$eIface.Name}}.{{end}}{{range $method := $eIface.Methods}}{{$outArgType := javaOutArgType $method false false $data.Imports $data.Env}}
		@Override
		public {{$outArgType}} {{toCamelCase $method.Name}}({{javaInArgs $method false false $data.Imports $data.Env}}) throws {{javaName "com/veyron2/ipc/VeyronException" $data.Imports}} {
			{{if ne $outArgType "void"}}return {{end}}this.{{javaClassToVarName $eIfaceClassName}}.{{toCamelCase $method.Name}}({{range $idx, $arg := $method.InArgs}}{{if gt $idx 0}}, {{end}}{{$arg.Name}}{{end}}{{if gt (len $method.InArgs) 0}}, {{end}}opts);
		}{{end}}
		{{end}}
	}{{end}}{{end}}
}
`

// javaUserImports stores Java class imports found in user-specified code (i.e.,
// structs, interfaces, consts).  These imports are stores as a map of
// ClassName -> ClassPkgPath.
type javaUserImports map[string]string

// javaConstImports returns Java class imports originating in the constant
// definitions of all VDL files in the provided package.
func javaConstImports(pkg *compile.Package, env *compile.Env) javaUserImports {
	var classPaths []string
	for _, file := range pkg.Files {
		for _, cdef := range file.ConstDefs {
			classPaths = append(classPaths, javaClassPaths(cdef.Value.Type(), false, env)...)
		}
	}
	return genJavaUserImports(classPaths, pkg)
}

// javaStructImports returns Java class imports originating in the provided
// struct definition.
func javaStructImports(tdef *compile.TypeDef, env *compile.Env) javaUserImports {
	classPaths := javaClassPaths(tdef.Type, false, env)
	for _, field := range javaStructFields(tdef.Type) {
		classPaths = append(classPaths, javaClassPaths(field.Type, false, env)...)
	}
	return genJavaUserImports(classPaths, tdef.File.Package)
}

// javaInterfaceImports returns Java class imports originating in the provided
// interface definition.
func javaInterfaceImports(iface *compile.Interface, isService bool, env *compile.Env) javaUserImports {
	classPaths := javaIfaceClassPaths(iface, isService, env)
	return genJavaUserImports(classPaths, iface.File.Package)
}

// javaClientsImports returns Java class imports for the Java client file, which
// provides client stub implementations for all the interfaces in the provided
// package.
func javaClientImports(pkg *compile.Package, env *compile.Env) javaUserImports {
	var classPaths []string
	for _, file := range pkg.Files {
		for _, iface := range file.Interfaces {
			classPaths = append(classPaths, javaIfaceClassPaths(iface, false, env)...)
			for _, eIface := range iface.Embeds {
				// Name of the embedded interface will be added twice, but that's fine.
				classPaths = append(classPaths, javaIfaceClassPaths(eIface, false, env)...)
			}
		}
	}
	// Add all the classes used in the client stub implementations.
	classPaths = append(classPaths, []string{
		path.Join(javaGenPkgPath(pkg.Path), "Client"),
		"com/veyron2/ipc/Client",
		"com/veyron2/runtime/RuntimeFactory",
		"java/lang/String",
		"java/lang/Object",
		"com/google/common/reflect/TypeToken",
	}...)
	return genJavaUserImports(classPaths, pkg)
}

// genJavaUserImports generates the final list of Java class imports, given the
// list of class paths.
func genJavaUserImports(classPaths []string, pkg *compile.Package) (ret javaUserImports) {
	// Sort all the class paths so that we get deterministic ordering;
	// it would be annoying for imported/non-imported class names to
	// change between runs.
	sort.StringSlice(classPaths).Sort()

	// Decide which class names to keep.  Here is the list of rules in order of
	// their preference:
	//   1) Prefer local classes.
	//   2) Prefer system classes (e.g., "java.lang.*").
	//   3) Prefer classes that are lexicographically smaller.
	ret = make(javaUserImports)
	for _, class := range classPaths {
		cur, ok := ret[path.Base(class)]
		switch {
		case !ok:
			ret[path.Base(class)] = path.Dir(class)
		case path.Dir(class) == pkg.Path:
			ret[path.Base(class)] = path.Dir(class)
		case path.Dir(class) == "java/lang" && path.Dir(cur) != pkg.Path:
			ret[path.Base(class)] = path.Dir(class)
		}
	}
	return
}

// javaClassPaths returns pathnames of all Java classes originating in the
// provided type.
func javaClassPaths(t *val.Type, forceClass bool, env *compile.Env) (ret []string) {
	if def := env.FindTypeDef(t); def != nil {
		return javaNamedClassPaths(def, forceClass)
	}
	switch t.Kind() {
	case val.List:
		ret = append(ret, "java/util/ArrayList")
		ret = append(ret, javaClassPaths(t.Elem(), true, env)...)
	case val.Map:
		ret = append(ret, "java/util/HashMap")
		ret = append(ret, javaClassPaths(t.Key(), true, env)...)
		ret = append(ret, javaClassPaths(t.Elem(), true, env)...)
	default:
		log.Fatalf("vdl: unhandled type %#v", t)
	}
	return
}

func javaNamedClassPaths(def *compile.TypeDef, forceClass bool) []string {
	if def.File == compile.GlobalFile {
		t, class := javaPrimitiveType(def, forceClass)
		if class {
			return []string{t}
		}
		return nil
	}
	return []string{path.Join(javaGenPkgPath(def.File.Package.Path), def.Name)}
}

// javaStreamClassPaths returns all class paths originating in the provided
// methods streaming argument.
func javaStreamClassPaths(m *compile.Method, env *compile.Env) (ret []string) {
	if !javaIsStreamingMethod(m) {
		return
	}
	ret = append(ret, "com/veyron2/vdl/Stream")
	if m.InStream != nil {
		ret = append(ret, javaClassPaths(m.InStream, true, env)...)
	} else {
		ret = append(ret, "java/lang/Void")
	}
	if m.OutStream != nil {
		ret = append(ret, javaClassPaths(m.OutStream, true, env)...)
	} else {
		ret = append(ret, "java/lang/Void")
	}
	return
}

// javaIfaceClassPaths returns all class paths originating in the provided
// interface.
func javaIfaceClassPaths(iface *compile.Interface, isService bool, env *compile.Env) (ret []string) {
	ret = []string{path.Join(javaGenPkgPath(iface.File.Package.Path), javaIfaceName(iface, isService))}
	for _, m := range iface.Methods {
		if javaIsStreamingMethod(m) {
			ret = append(ret, javaStreamClassPaths(m, env)...)
		}
		if !isService {
			ret = append(ret, "com/veyron2/ipc/Client")
		} else {
			ret = append(ret, "com/veyron2/ipc/Context")
		}
		for _, iarg := range m.InArgs {
			ret = append(ret, javaClassPaths(iarg.Type, false, env)...)
		}
		for _, oarg := range m.OutArgs {
			ret = append(ret, javaClassPaths(oarg.Type, false, env)...)
		}
		ret = append(ret, "com/veyron2/ipc/VeyronException")
	}
	for _, eIface := range iface.Embeds {
		ret = append(ret, path.Join(javaGenPkgPath(eIface.File.Package.Path), javaIfaceName(eIface, isService)))
	}
	return
}

// javaType returns the Java type string given the provided VDL type.  It
// consults the provided imports map to see if the full pathnames or only
// base names can be used (e.g., "java.util.HashMap" or "HashMap").
func javaType(t *val.Type, forceClass bool, imports javaUserImports, env *compile.Env) string {
	if def := env.FindTypeDef(t); def != nil {
		return javaNamedType(def, forceClass, imports)
	}
	switch t.Kind() {
	case val.List:
		return fmt.Sprintf("%s<%s>", javaName("java/util/ArrayList", imports), javaType(t.Elem(), true, imports, env))
	case val.Map:
		return fmt.Sprintf("%s<%s, %s>", javaName("java/util/HashMap", imports), javaType(t.Key(), true, imports, env), javaType(t.Elem(), true, imports, env))
	default:
		log.Fatalf("vdl: unhandled type %v %v", t.Kind(), t)
		return ""
	}
}

func javaNamedType(def *compile.TypeDef, forceClass bool, imports javaUserImports) string {
	if def.File == compile.GlobalFile {
		name, class := javaPrimitiveType(def, forceClass)
		if class {
			return javaName(name, imports)
		}
		return name
	}
	return javaQualifiedName(def.Name, def.File, imports)
}

func javaPrimitiveType(def *compile.TypeDef, forceClass bool) (string, bool) {
	if def.Name == "error" {
		return "com/veyron2/ipc/VeyronException", true
	}
	switch def.Type.Kind() {
	case val.Bool:
		if forceClass {
			return "java/lang/Boolean", true
		} else {
			return "boolean", false
		}
	case val.Uint32, val.Int32:
		if forceClass {
			return "java/lang/Integer", true
		} else {
			return "int", false
		}
	case val.Uint64, val.Int64:
		if forceClass {
			return "java/lang/Long", true
		} else {
			return "long", false
		}
	case val.Float32:
		if forceClass {
			return "java/lang/Float", true
		} else {
			return "float", false
		}
	case val.Float64:
		if forceClass {
			return "java/lang/Double", true
		} else {
			return "double", false
		}
	// TODO(spetrovic) handle val.Complex types.
	case val.String:
		return "java/lang/String", true
	case val.Any:
		return "java/lang/Object", true
	default:
		log.Fatalf("unexpected primitive type %q", def)
		return "", false
	}
}

func javaStreamType(method *compile.Method, imports javaUserImports, env *compile.Env) string {
	if !javaIsStreamingMethod(method) {
		return ""
	}
	stream := javaName("com/veyron2/vdl/Stream", imports) + "<"
	if method.InStream != nil {
		stream += javaType(method.InStream, true, imports, env)
	} else {
		stream += javaName("java/lang/Void", imports)
	}
	stream += ", "
	if method.OutStream != nil {
		stream += javaType(method.OutStream, true, imports, env)
	} else {
		stream += javaName("java/lang/Void", imports)
	}
	stream += ">"
	return stream
}

// javaConstVal returns the value string for the provided constant.
func javaConstVal(v *val.Value) string {
	switch v.Kind() {
	case val.Bool:
		if v.Bool() {
			return "true"
		} else {
			return "false"
		}
	case val.Int32, val.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case val.Uint32, val.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case val.Float32, val.Float64:
		return strconv.FormatFloat(v.Float(), 'g', -1, bitlen(v.Kind()))
	// TODO(spetrovic): handle complex numbers.
	case val.String:
		return strconv.Quote(v.String())
	case val.Bytes:
		return strconv.Quote(string(v.Bytes()))
	}
	// TODO(spetrovic): Handle Enum, List, Map, Struct, OneOf, Any
	panic(fmt.Errorf("vdl: valueGo unhandled type %v %v", v.Kind(), v.Type()))
}

// javaImports prunes the provided imports list, removing the unnecessary
// imports, and returns the pruned list in the Java format.
func javaImports(imports javaUserImports, pkg *compile.Package) (ret []string) {
	for base, dir := range imports {
		if dir == javaGenPkgPath(pkg.Path) || dir == "java/lang" {
			continue
		}
		ret = append(ret, javaPath(path.Join(dir, base)))
	}
	sort.StringSlice(ret).Sort()
	return
}

// javaNonErrorOutArgs returns the list of non-error output arguments for the
// provided method.
func javaNonErrorOutArgs(method *compile.Method) []*compile.Arg {
	return method.OutArgs[:len(method.OutArgs)-1]
}

// javaOutArg returns the type for the (single) output argument for the provided
// method.
func javaOutArgType(method *compile.Method, isService, forceClass bool, imports javaUserImports, env *compile.Env) string {
	if javaIsStreamingMethod(method) {
		if isService {
			// Services have the stream inlined with the input arguments.
			return "void"
		}
		// Clients get a single output stream argument.
		return javaStreamType(method, imports, env)

		// TODO(spetrovic): All the other output arguments associated with this
		// method are moved onto the stream.  Implement this functionality.
	}
	switch len(method.OutArgs) {
	case 0:
		log.Fatalf("vdl: method must have at least one output arg (i.e., error): %v", method)
		return ""
	case 1:
		return "void"
	case 2:
		return javaType(method.OutArgs[0].Type, forceClass, imports, env)
	default:
		return method.Name + "Out"
	}
}

func javaOutArgs(method *compile.Method) []*compile.Arg {
	// Ignore the last (i.e., error) argument.
	return method.OutArgs[:len(method.OutArgs)-1]
}

func javaIsStreamingMethod(method *compile.Method) bool {
	return method.InStream != nil || method.OutStream != nil
}

func javaShouldPackOutArgs(method *compile.Method, isService bool, imports javaUserImports, env *compile.Env) bool {
	return javaOutArgType(method, isService, false, imports, env) == (method.Name + "Out")
}

// javaInArgs returns the in-args string for the provided method.
func javaInArgs(method *compile.Method, isService, camelCase bool, imports javaUserImports, env *compile.Env) (ret string) {
	if isService {
		ret += javaName("com/veyron2/ipc/Context", imports) + " context"
	}
	for _, arg := range method.InArgs {
		if ret != "" {
			ret += ", "
		}
		argName := arg.Name
		if camelCase {
			argName = toCamelCase(argName)
		}
		ret += javaType(arg.Type, false, imports, env) + " " + argName
	}
	if !isService {
		if ret != "" {
			ret += ", "
		}
		ret += javaName("com/veyron2/ipc/Client", imports) + ".CallOption... opts"
	}
	if javaIsStreamingMethod(method) && isService {
		// Services get the stream inlined with the input arguments; clients
		// treat the stream as an output argument so we ignore it here.
		if ret != "" {
			ret += ", "
		}
		ret += javaStreamType(method, imports, env) + " stream"
	}
	return
}

// javaAllEmbeddedIfaces returns all unique interfaces in the embed tree
// starting at the provided interface (not including that interface).
func javaAllEmbeddedIfaces(iface *compile.Interface) (ret []*compile.Interface) {
	added := make(map[string]bool)
	for _, eIface := range iface.Embeds {
		for _, eIface = range append(javaAllEmbeddedIfaces(eIface), eIface) {
			path := path.Join(eIface.File.Package.Path, eIface.Name)
			if _, ok := added[path]; ok { // already added iface
				continue
			}
			ret = append(ret, eIface)
			added[path] = true
		}
	}
	return
}

// javaStructFields returns the types of all the fields in the provided struct type.
func javaStructFields(t *val.Type) (ret []val.StructField) {
	for idx := 0; idx < t.NumField(); idx++ {
		ret = append(ret, t.Field(idx))
	}
	return
}

// javaDoc transforms the provided VDL comment into the JavaDoc format.
func javaDoc(goComment string) string {
	if goComment == "" {
		return ""
	}
	return "/**\n" + strings.Replace(goComment, "//", " *", -1) + "**/\n"
}

// javaPath converts the provided Go path into the Java path.  It replaces all "/"
// with "." in the path.
func javaPath(goPath string) string {
	return strings.Replace(goPath, "/", ".", -1)
}

// javaName converts the provided Go pathname into the Java pathname and returns a
// shorter version of it if possible, given the provided imports.
func javaName(goPathName string, imports javaUserImports) string {
	if dir, ok := imports[path.Base(goPathName)]; ok && dir == path.Dir(goPathName) {
		return path.Base(goPathName)
	}
	return javaPath(goPathName)
}

// javaName returns the fully qualified Java name for the provided goName.
func javaQualifiedName(goName string, file *compile.File, imports javaUserImports) string {
	return javaName(path.Join(javaGenPkgPath(file.Package.Path), goName), imports)
}

// javaIfaceName returns the name for the provided interface.
func javaIfaceName(iface *compile.Interface, isService bool) string {
	if !isService {
		return iface.Name
	}
	return iface.Name + "Service"
}

// javaClassToVarName generates a Java variable name corresponding to the provided
// Java class name.
func javaClassToVarName(className string) string {
	dir := path.Dir(className)
	if dir == "." {
		return toCamelCase(className)
	}
	return strings.Replace(path.Dir(className), "/", "_", -1) + "_" + toCamelCase(path.Base(className))
}

// toCamelCase converts ThisString to thisString.
func toCamelCase(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}

// toConstCase converts ThisString to THIS_STRING.
func toConstCase(s string) string {
	var buf bytes.Buffer
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			buf.WriteRune('_')
		}
		buf.WriteRune(unicode.ToUpper(r))
	}
	return buf.String()
}

var (
	javaPkgTmpl    *template.Template
	javaConstTmpl  *template.Template
	javaClassTmpl  *template.Template
	javaIfaceTmpl  *template.Template
	javaClientTmpl *template.Template
)

// The template mechanism is great at high-level formatting and simple
// substitution, but is bad at more complicated logic (by design).  We define
// some functions that we can use in the template so that when things get
// complicated we back off to a regular function.
func init() {
	funcMap := template.FuncMap{
		"javaDoc":               javaDoc,
		"javaGenPkgPath":        javaGenPkgPath,
		"javaPath":              javaPath,
		"javaImports":           javaImports,
		"javaType":              javaType,
		"javaConstVal":          javaConstVal,
		"javaInArgs":            javaInArgs,
		"javaOutArgType":        javaOutArgType,
		"javaShouldPackOutArgs": javaShouldPackOutArgs,
		"javaNonErrorOutArgs":   javaNonErrorOutArgs,
		"javaIsStreamingMethod": javaIsStreamingMethod,
		"javaOutArgs":           javaOutArgs,
		"javaStructFields":      javaStructFields,
		"javaAllEmbeddedIfaces": javaAllEmbeddedIfaces,
		"javaName":              javaName,
		"javaQualifiedName":     javaQualifiedName,
		"javaIfaceName":         javaIfaceName,
		"javaClassToVarName":    javaClassToVarName,
		"toCamelCase":           toCamelCase,
		"toConstCase":           toConstCase,
	}
	javaPkgTmpl = template.Must(template.New("javaPkg").Funcs(funcMap).Parse(javaPkgTmplStr))
	javaConstTmpl = template.Must(template.New("javaConst").Funcs(funcMap).Parse(javaConstTmplStr))
	javaClassTmpl = template.Must(template.New("javaClass").Funcs(funcMap).Parse(javaClassTmplStr))
	javaIfaceTmpl = template.Must(template.New("javaIface").Funcs(funcMap).Parse(javaIfaceTmplStr))
	javaClientTmpl = template.Must(template.New("javaClient").Funcs(funcMap).Parse(javaClientTmplStr))
}
