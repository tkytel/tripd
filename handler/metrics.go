package handler

import (
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/tkytel/tripd/utils"
)

func HandleMetrics(c *fiber.Ctx) error {
	return c.SendString(GenerateMetrics())
}

func GenerateMetrics() string {
	about, err := GenerateAbout()
	if err != nil {
		return ""
	}
	res := ""

	// tripd_about
	v := reflect.ValueOf(about)
	t := reflect.TypeOf(about)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}

	for i := range v.NumField() {
		field := t.Field(i)
		value := v.Field(i).Interface()
		json, ok_json := field.Tag.Lookup("json")
		desc, ok_desc := field.Tag.Lookup("desc")

		if ok_json && ok_desc {
			res += fmt.Sprintf(
				"# HELP tripd_about_%v %v\n",
				json,
				desc,
			)
			res += fmt.Sprintf(
				"# TYPE tripd_about_%v gauge\n",
				json,
			)
			res += fmt.Sprintf(
				"tripd_about_%v{identifier=\"%v\",value=\"%v\"} 1\n",
				json,
				about.Identifier,
				value,
			)
		}
	}

	res += "# HELP tripd_peer peering statistics seen from this PBX system"
	res += "# TYPE tripd_peer gauge"

	for _, vp := range utils.Peers {
		loss := (float64)(0)
		if vp.Loss != nil {
			loss = *vp.Loss
		}
		rtt := (float64)(0)
		if vp.Loss != nil {
			rtt = *vp.Rtt
		}

		res += fmt.Sprintf(
			"tripd_peer{identifier=\"%v\",peer=\"%v\",key=\"loss\"} %v\n",
			about.Identifier,
			vp.Identifier,
			loss,
		)
		res += fmt.Sprintf(
			"tripd_peer{identifier=\"%v\",peer=\"%v\",key=\"rtt\"} %v\n",
			about.Identifier,
			vp.Identifier,
			rtt,
		)

		measurable := 0
		if vp.Measurable {
			measurable = 1
		}
		res += fmt.Sprintf(
			"tripd_peer{identifier=\"%v\",peer=\"%v\",key=\"measurable\"} %v\n",
			about.Identifier,
			vp.Identifier,
			measurable,
		)
	}

	return res
}
