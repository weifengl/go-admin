package controller

import (
	c "github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
)

var config c.Config

func SetConfig(cfg c.Config) {
	config = cfg
}

func aAlert() types.AlertAttribute {
	return aTemplate().Alert()
}

func aForm() types.FormAttribute {
	return aTemplate().Form()
}

func aRow() types.RowAttribute {
	return aTemplate().Row()
}

func aCol() types.ColAttribute {
	return aTemplate().Col()
}

func aTree() types.TreeAttribute {
	return aTemplate().Tree()
}

func aDataTable() types.DataTableAttribute {
	return aTemplate().DataTable()
}

func aTable() types.TableAttribute {
	return aTemplate().Table()
}

func aBox() types.BoxAttribute {
	return aTemplate().Box()
}

func aTemplate() template.Template {
	return template.Get(config.THEME)
}
