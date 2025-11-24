package model

type ModuleDef struct {
	Name        string
	Description string
	Meta        ModuleDefMeta
}

type ModuleDefMeta struct {
	Permissions []ModuleDefMetaPermission
}

type ModuleDefMetaPermission struct {
	Scope string
	Title string
}
