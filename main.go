package main

import (
	"action"
	"system"

	"fieldtype"
)

func main() {
	var merp = system.NewApplication()

	// geo
	var geo = merp.NewModule("Geo", "Merp", "../system")

	var cGeoLocationsList = geo.NewGenericController("LocationsList", false)
	var mGeoLocationsList_id = system.NewField("_id", fieldtype.INT, -1, "", false, true, true, "Identificador", nil)
	var mGeoLocationsList_location = system.NewField("_town", fieldtype.VARCHAR, 512, "", false, false, false, "Localizaci&oacute;n", nil)
	var mGeoLocationsList = cGeoLocationsList.NewGenericModel(action.SET, []*system.Field{
		mGeoLocationsList_id,
		mGeoLocationsList_location,
	})

	var cGeoDependences = geo.NewGenericController("Dependences", false)
	var mGeoDependences_id = system.NewField("_id", fieldtype.VARCHAR, 4, "0", false, true, false, "Identificador", nil)
	var mGeoDependences_description = system.NewField("_description", fieldtype.VARCHAR, 32, "", false, false, false, "Descripci&oacute;n", nil)
	var mGeoDependences = cGeoDependences.NewGenericModel(action.SET, []*system.Field{
		mGeoDependences_id,
		mGeoDependences_description,
	})

	// geo:enrollmentmodes
	var cGeoEnrollmentmodes = geo.NewGenericController("Enrollmentmodes", false)
	var mGeoEnrollmentmodes_id = system.NewField("_id", fieldtype.VARCHAR, 8, "", false, true, false, "Identificador", nil)
	var mGeoEnrollmentmodes_description = system.NewField("_description", fieldtype.VARCHAR, 64, "", false, false, false, "Descripci&oacute;n", nil)
	var mGeoEnrollmentmodes = cGeoEnrollmentmodes.NewGenericModel(action.SET, []*system.Field{
		mGeoEnrollmentmodes_id,
		mGeoEnrollmentmodes_description,
	})

	// geo:typologies
	var cGeoTypologies = geo.NewGenericController("Typologies", false)
	var mGeoTypologies_id = system.NewField("_id", fieldtype.INT, -1, "", false, true, true, "Identificador", nil)
	var mGeoTypologies_description = system.NewField("_description", fieldtype.VARCHAR, 32, "", false, false, false, "Descripci&oacute;n", nil)
	/* var mGeoTypologies =*/ cGeoTypologies.NewGenericModel(action.SET, []*system.Field{
		mGeoTypologies_id,
		mGeoTypologies_description,
	})

	var cGeoSchoolsEnum = geo.NewGenericController("SchoolsEnum", false)
	var mGeoSchoolsEnum_id = system.NewField("_id", fieldtype.INT, -1, "", false, true, true, "Identificador", nil)
	var mGeoSchoolsEnum_name = system.NewField("_name", fieldtype.VARCHAR, 128, "", false, false, false, "Instituci&oacute;n", nil)
	var mGeoSchoolsEnum = cGeoSchoolsEnum.NewGenericModel(action.SET, []*system.Field{
		mGeoSchoolsEnum_id,
		mGeoSchoolsEnum_name,
	})

	var cGeoCantonsEnum = geo.NewGenericController("CantonsEnum", false)
	var mGeoCantonsEnum_id = system.NewField("_id", fieldtype.INT, -1, "-1", false, true, true, "Identificador", nil)
	var mGeoCantonsEnum_name = system.NewField("_name", fieldtype.VARCHAR, 64, "", false, false, false, "Nombre", nil)
	var mGeoCantonsEnum = cGeoCantonsEnum.NewGenericModel(action.SET, []*system.Field{
		mGeoCantonsEnum_id,
		mGeoCantonsEnum_name,
	})

	var cGeoDistrictsEnum = geo.NewGenericController("DistrictsEnum", false)
	var mGeoDistrictsEnum_id = system.NewField("_id", fieldtype.INT, -1, "-1", false, true, true, "Identificador", nil)
	var mGeoDistrictsEnum_name = system.NewField("_name", fieldtype.VARCHAR, 128, "", false, false, false, "Nombre", nil)
	var mGeoDistrictsEnum = cGeoDistrictsEnum.NewGenericModel(action.SET, []*system.Field{
		mGeoDistrictsEnum_id,
		mGeoDistrictsEnum_name,
	})
	/* -------------------------------------------------------------------------------------------------------------------------------- */
	/*
		geo.AddModel(system.NewGenericModel("GeoProvinces", "Merp\\Geo", false, []*system.Field{
			system.NewField("_id", fieldtype.INT, -1, "0", false, true, false, "Identificador", nil),
			system.NewField("_name", fieldtype.VARCHAR, 32, "0", true, false, false, "Nombre", nil),
			system.NewField("_area", fieldtype.DECIMAL, 11, "0", true, false, false, "&Aacute;rea (km&#178;)", nil),
		}))
	*/

	// geo:cantonsdetail
	var cGeoCantons = geo.NewGenericController("Cantonsdetail", true)
	cGeoCantons.NewGenericModel(action.GET, []*system.Field{
		system.NewField("_id", fieldtype.INT, -1, "-1", false, true, true, "Identificador", nil),
		system.NewField("_cantonid", fieldtype.INT, -1, "-1", false, false, false, "C&oacute;digo", nil),
		system.NewField("_cantonname", fieldtype.VARCHAR, 64, "", false, false, false, "Cant&oacute;n", nil),
		system.NewField("_cantonarea", fieldtype.DECIMAL, 16, "", false, false, false, "&Aacute;rea (km&#178;)", nil),
		system.NewField("_ids", fieldtype.DECIMAL, 8, "", false, false, false, "IDS", nil),
		system.NewField("_age", fieldtype.YEAR, -1, "", false, false, false, "A&ntilde;o", nil),
	})

	cGeoCantons.NewGenericModel(action.SET, []*system.Field{
		system.NewField("_id", fieldtype.INT, -1, "", false, true, true, "Identificador", nil),
		system.NewField("_canton", fieldtype.INT, -1, "", false, false, false, "Cant&oacute;n", system.NewEnum(mGeoCantonsEnum, mGeoCantonsEnum_id, mGeoCantonsEnum_name)),
		system.NewField("_ids", fieldtype.DECIMAL, 8, "", false, false, false, "IDS", nil),
		system.NewField("_year", fieldtype.YEAR, -1, "", false, false, false, "A&ntilde;o", nil),
	})

	// geo:districtsdetail
	var cGeoDistricts = geo.NewGenericController("Districtsdetail", true)
	cGeoDistricts.NewGenericModel(action.GET, []*system.Field{
		system.NewField("_id", fieldtype.INT, -1, "-1", false, true, true, "Identificador", nil),
		system.NewField("_districtid", fieldtype.INT, -1, "-1", false, false, false, "C&oacute;digo", nil),
		system.NewField("_districtname", fieldtype.VARCHAR, 128, "", false, false, false, "Distrito", nil),
		system.NewField("_districtarea", fieldtype.DECIMAL, 16, "", false, false, false, "&Aacute;rea (km&#178;)", nil),
		system.NewField("_population", fieldtype.DECIMAL, 8, "", false, false, false, "Poblaci&oacute;n", nil),
		system.NewField("_economic", fieldtype.DECIMAL, 8, "", false, false, false, "Econ&oacute;mica", nil),
		system.NewField("_voterturnout", fieldtype.DECIMAL, 8, "", false, false, false, "Electoral", nil),
		system.NewField("_health", fieldtype.DECIMAL, 8, "", false, false, false, "Salud", nil),
		system.NewField("_education", fieldtype.DECIMAL, 8, "", false, false, false, "Educaci&oacute;n", nil),
		system.NewField("_security", fieldtype.DECIMAL, 8, "", false, false, false, "Seguridad", nil),
		system.NewField("_ids", fieldtype.DECIMAL, 8, "", false, false, false, "IDS", nil),
		system.NewField("_age", fieldtype.YEAR, -1, "", false, false, false, "A&ntilde;o", nil),
	})

	cGeoDistricts.NewGenericModel(action.SET, []*system.Field{
		system.NewField("_id", fieldtype.INT, -1, "", false, true, true, "Identificador", nil),
		system.NewField("_district", fieldtype.INT, -1, "", false, false, false, "Distrito", system.NewEnum(mGeoDistrictsEnum, mGeoDistrictsEnum_id, mGeoDistrictsEnum_name)),
		system.NewField("_population", fieldtype.DECIMAL, 8, "", false, false, false, "Poblaci&oacute;n", nil),
		system.NewField("_economic", fieldtype.DECIMAL, 8, "", false, false, false, "Econ&oacute;mica", nil),
		system.NewField("_voterturnout", fieldtype.DECIMAL, 8, "", false, false, false, "Electoral", nil),
		system.NewField("_health", fieldtype.DECIMAL, 8, "", false, false, false, "Salud", nil),
		system.NewField("_education", fieldtype.DECIMAL, 8, "", false, false, false, "Educaci&oacute;n", nil),
		system.NewField("_security", fieldtype.DECIMAL, 8, "", false, false, false, "Seguridad", nil),
		system.NewField("_ids", fieldtype.DECIMAL, 8, "", false, false, false, "IDS", nil),
		system.NewField("_year", fieldtype.YEAR, -1, "", false, false, false, "A&ntilde;o", nil),
	})

	// geo:regionals
	var cGeoRegionals = geo.NewGenericController("Regionals", false)
	var mGeoRegionals_id = system.NewField("_id", fieldtype.INT, -1, "", false, true, true, "Identificador", nil)
	var mGeoRegionals_name = system.NewField("_name", fieldtype.VARCHAR, 32, "", false, false, false, "Direcci&oacute;n regional", nil)
	var fGeoRegionals = []*system.Field{
		mGeoRegionals_id,
		mGeoRegionals_name,
	}

	var mGeoRegionals = cGeoRegionals.NewGenericModel(action.SET, fGeoRegionals)

	// geo:schools
	var cGeoSchools = geo.NewGenericController("Schools", false)
	cGeoSchools.NewGenericModel(action.GET, []*system.Field{
		//                nombre, tipo, tamaño, valor por defecto, permite nulos?, es llave primaria?, es autoincrementable?, descripción, opciones
		system.NewField("_id", fieldtype.INT, -1, "", false, true, true, "Identificador", nil),
		system.NewField("_circuit", fieldtype.VARCHAR, 2, "00", false, false, false, "Circuito", nil),
		system.NewField("_town", fieldtype.VARCHAR, 512, "", false, false, false, "Localizaci&oacute;n", nil),
		system.NewField("_regional", fieldtype.VARCHAR, 32, "", false, false, false, "Regional", nil),
		system.NewField("_budgetcode", fieldtype.VARCHAR, 4, "0000", true, false, false, "C&oacute;digo", nil),
		system.NewField("_name", fieldtype.VARCHAR, 64, "", false, false, false, "Instituci&oacute;n", nil),
		system.NewField("_dependence", fieldtype.VARCHAR, 32, "", false, false, false, "Dependencia", nil),
		system.NewField("_phone", fieldtype.TEL, -1, "", false, false, false, "Tel&eacute;fono", nil),
	})

	cGeoSchools.NewGenericModel(action.SET, []*system.Field{
		system.NewField("_id", fieldtype.INT, -1, "", false, true, true, "Identificador", nil),
		system.NewField("_circuit", fieldtype.VARCHAR, 2, "00", false, false, false, "Circuito", nil),
		system.NewField("_town", fieldtype.INT, -1, "0", false, false, false, "Poblado", system.NewEnum(mGeoLocationsList, mGeoLocationsList_id, mGeoLocationsList_location)),
		system.NewField("_regional", fieldtype.INT, -1, "0", false, false, false, "Direcci&oacute;n regional", system.NewEnum(mGeoRegionals, mGeoRegionals_id, mGeoRegionals_name)),
		system.NewField("_budgetcode", fieldtype.VARCHAR, 4, "", true, false, false, "C&oacute;digo presupuestario", nil),
		system.NewField("_name", fieldtype.VARCHAR, 64, "", false, false, false, "Instituci&oacute;n", nil),
		system.NewField("_dependence", fieldtype.VARCHAR, 4, "0", false, false, false, "Dependencia", system.NewEnum(mGeoDependences, mGeoDependences_id, mGeoDependences_description)),
		system.NewField("_phone", fieldtype.TEL, -1, "", false, false, false, "Tel&eacute;fono", nil),
	})

	// geo:enrollments
	var cGeoEnrollments = geo.NewGenericController("Enrollments", true)

	cGeoEnrollments.NewGenericModel(action.GET, []*system.Field{
		system.NewField("_id", fieldtype.INT, -1, "-1", false, true, true, "ID", nil),
		system.NewField("_regional", fieldtype.VARCHAR, 32, "", false, false, false, "Regional", nil),
		system.NewField("_circuit", fieldtype.VARCHAR, 2, "00", false, false, false, "Circuito", nil),
		system.NewField("_budgetcode", fieldtype.VARCHAR, 8, "", true, false, false, "C&oacute;digo", nil),
		system.NewField("_school", fieldtype.VARCHAR, 64, "0000", false, false, false, "Instituci&oacute;n", nil),
		system.NewField("_town", fieldtype.VARCHAR, 256, "", false, false, false, "Poblado", nil),
		system.NewField("_dependence", fieldtype.VARCHAR, 32, "", false, false, false, "Dependencia", nil),
		system.NewField("_zone", fieldtype.VARCHAR, 32, "", false, false, false, "Zona", nil),
		system.NewField("_phone", fieldtype.TEL, -1, "", false, false, false, "Tel&eacute;fono", nil),

		system.NewField("_enrollmentmode", fieldtype.VARCHAR, 64, "", false, false, false, "Modalidad", nil),
		system.NewField("_amount", fieldtype.INT, -1, "0", false, false, false, "Matr&iacute;cula", nil),
	})

	cGeoEnrollments.NewGenericModel(action.SET, []*system.Field{
		system.NewField("_id", fieldtype.INT, -1, "", false, true, true, "Identificador", nil),
		system.NewField("_school", fieldtype.VARCHAR, 64, "", false, false, false, "Instituci&oacute;n", system.NewEnum(mGeoSchoolsEnum, mGeoSchoolsEnum_id, mGeoSchoolsEnum_name)),
		system.NewField("_enrollmentmode", fieldtype.VARCHAR, 8, "", false, false, false, "Modalidad", system.NewEnum(mGeoEnrollmentmodes, mGeoEnrollmentmodes_id, mGeoEnrollmentmodes_description)),
		system.NewField("_year", fieldtype.YEAR, -1, "", false, false, false, "A&ntilde;o", nil),
		system.NewField("_amount", fieldtype.INT, -1, "0", false, false, false, "Matr&iacute;cula", nil),
	})

	// geo:enrollments1
	var cGeoEnrollments1 = geo.NewGenericController("Enrollments1", true)

	cGeoEnrollments1.NewGenericModel(action.GET, []*system.Field{
		system.NewField("_id", fieldtype.INT, -1, "-1", false, true, true, "ID", nil),
		system.NewField("_regional", fieldtype.VARCHAR, 32, "", false, false, false, "Regional", nil),
		system.NewField("_circuit", fieldtype.VARCHAR, 2, "00", false, false, false, "Circuito", nil),
		system.NewField("_budgetcode", fieldtype.VARCHAR, 4, "", true, false, false, "C&oacute;digo", nil),
		system.NewField("_name", fieldtype.VARCHAR, 64, "0000", false, false, false, "Instituci&oacute;n", nil),
		system.NewField("_town", fieldtype.VARCHAR, 256, "", false, false, false, "Poblado", nil),
		system.NewField("_dependence", fieldtype.VARCHAR, 32, "", false, false, false, "Dependencia", nil),
		system.NewField("_zone", fieldtype.VARCHAR, 32, "", false, false, false, "Zona", nil),
		system.NewField("_phone", fieldtype.TEL, -1, "", false, false, false, "Tel&eacute;fono", nil),

		system.NewField("_rc", fieldtype.INT, -1, "0", false, false, false, "RC", nil),
		system.NewField("_pe", fieldtype.INT, -1, "0", false, false, false, "PE", nil),
	})

	// geo:enrollment2
	var cGeoEnrollments2 = geo.NewGenericController("Enrollments2", true)

	cGeoEnrollments2.NewGenericModel(action.GET, []*system.Field{
		system.NewField("_id", fieldtype.INT, -1, "-1", false, true, true, "ID", nil),
		system.NewField("_regional", fieldtype.VARCHAR, 32, "", false, false, false, "Regional", nil),
		system.NewField("_circuit", fieldtype.VARCHAR, 2, "00", false, false, false, "Circuito", nil),
		system.NewField("_budgetcode", fieldtype.VARCHAR, 4, "", true, false, false, "C&oacute;digo", nil),
		system.NewField("_name", fieldtype.VARCHAR, 64, "0000", false, false, false, "Instituci&oacute;n", nil),
		system.NewField("_town", fieldtype.VARCHAR, 256, "", false, false, false, "Poblado", nil),
		system.NewField("_dependence", fieldtype.VARCHAR, 32, "", false, false, false, "Dependencia", nil),
		system.NewField("_zone", fieldtype.VARCHAR, 32, "", false, false, false, "Zona", nil),
		system.NewField("_phone", fieldtype.TEL, -1, "", false, false, false, "Tel&eacute;fono", nil),

		system.NewField("_ad", fieldtype.INT, -1, "0", false, false, false, "AD", nil),
		system.NewField("_ae", fieldtype.INT, -1, "0", false, false, false, "AE", nil),
		system.NewField("_pr", fieldtype.INT, -1, "0", false, false, false, "PR", nil),
		system.NewField("_rc", fieldtype.INT, -1, "0", false, false, false, "RC", nil),
		system.NewField("_pe", fieldtype.INT, -1, "0", false, false, false, "PE", nil),
		system.NewField("_pea", fieldtype.INT, -1, "0", false, false, false, "PEA", nil),
	})

	// geo:enrollments3
	var cGeoEnrollments3 = geo.NewGenericController("Enrollments3", true)

	cGeoEnrollments3.NewGenericModel(action.GET, []*system.Field{
		system.NewField("_id", fieldtype.INT, -1, "-1", false, true, true, "ID", nil),
		system.NewField("_regional", fieldtype.VARCHAR, 32, "", false, false, false, "Regional", nil),
		system.NewField("_circuit", fieldtype.VARCHAR, 2, "00", false, false, false, "Circuito", nil),
		system.NewField("_budgetcode", fieldtype.VARCHAR, 4, "", true, false, false, "C&oacute;digo", nil),
		system.NewField("_name", fieldtype.VARCHAR, 64, "0000", false, false, false, "Instituci&oacute;n", nil),
		system.NewField("_town", fieldtype.VARCHAR, 256, "", false, false, false, "Poblado", nil),
		system.NewField("_dependence", fieldtype.VARCHAR, 32, "", false, false, false, "Dependencia", nil),
		system.NewField("_zone", fieldtype.VARCHAR, 32, "", false, false, false, "Zona", nil),
		system.NewField("_phone", fieldtype.TEL, -1, "", false, false, false, "Tel&eacute;fono", nil),

		system.NewField("_bi", fieldtype.INT, -1, "0", false, false, false, "BI", nil),
		system.NewField("_se", fieldtype.INT, -1, "0", false, false, false, "SE", nil),
		system.NewField("_pn", fieldtype.INT, -1, "0", false, false, false, "PN", nil),
		system.NewField("_pea", fieldtype.INT, -1, "0", false, false, false, "PEA", nil),
	})

	// geo:enrollments4
	var cGeoEnrollments4 = geo.NewGenericController("Enrollments4", true)

	cGeoEnrollments4.NewGenericModel(action.GET, []*system.Field{
		system.NewField("_id", fieldtype.INT, -1, "-1", false, true, true, "ID", nil),
		system.NewField("_regional", fieldtype.VARCHAR, 32, "", false, false, false, "Regional", nil),
		system.NewField("_circuit", fieldtype.VARCHAR, 2, "00", false, false, false, "Circuito", nil),
		system.NewField("_budgetcode", fieldtype.VARCHAR, 4, "", true, false, false, "C&oacute;digo", nil),
		system.NewField("_name", fieldtype.VARCHAR, 64, "0000", false, false, false, "Instituci&oacute;n", nil),
		system.NewField("_town", fieldtype.VARCHAR, 256, "", false, false, false, "Poblado", nil),
		system.NewField("_dependence", fieldtype.VARCHAR, 32, "", false, false, false, "Dependencia", nil),
		system.NewField("_zone", fieldtype.VARCHAR, 32, "", false, false, false, "Zona", nil),
		system.NewField("_phone", fieldtype.TEL, -1, "", false, false, false, "Tel&eacute;fono", nil),

		system.NewField("_cnvmts", fieldtype.INT, -1, "0", false, false, false, "Matrícula", nil),
	})

	// geo:enrollments5
	var cGeoEnrollments5 = geo.NewGenericController("Enrollments5", true)

	cGeoEnrollments5.NewGenericModel(action.GET, []*system.Field{
		system.NewField("_id", fieldtype.INT, -1, "0", false, true, true, "ID", nil),
		system.NewField("_regional", fieldtype.VARCHAR, 32, "", false, false, false, "Regional", nil),
		system.NewField("_circuit", fieldtype.VARCHAR, 2, "00", false, false, false, "Circuito", nil),
		system.NewField("_budgetcode", fieldtype.VARCHAR, 4, "", true, false, false, "C&oacute;digo", nil),
		system.NewField("_name", fieldtype.VARCHAR, 64, "0000", false, false, false, "Instituci&oacute;n", nil),
		system.NewField("_town", fieldtype.VARCHAR, 256, "", false, false, false, "Poblado", nil),
		system.NewField("_dependence", fieldtype.VARCHAR, 32, "", false, false, false, "Dependencia", nil),
		system.NewField("_zone", fieldtype.VARCHAR, 32, "", false, false, false, "Zona", nil),
		system.NewField("_phone", fieldtype.TEL, -1, "", false, false, false, "Tel&eacute;fono", nil),

		system.NewField("_cee", fieldtype.INT, -1, "0", false, false, false, "Matrícula", nil),
	})

	// geo:enrollments6
	var cGeoEnrollments6 = geo.NewGenericController("Enrollments6", true)

	cGeoEnrollments6.NewGenericModel(action.GET, []*system.Field{
		system.NewField("_id", fieldtype.INT, -1, "-1", false, true, true, "ID", nil),
		system.NewField("_regional", fieldtype.VARCHAR, 32, "", false, false, false, "Regional", nil),
		system.NewField("_circuit", fieldtype.VARCHAR, 2, "00", false, false, false, "Circuito", nil),
		system.NewField("_budgetcode", fieldtype.VARCHAR, 4, "", true, false, false, "C&oacute;digo", nil),
		system.NewField("_name", fieldtype.VARCHAR, 64, "0000", false, false, false, "Instituci&oacute;n", nil),
		system.NewField("_town", fieldtype.VARCHAR, 256, "", false, false, false, "Poblado", nil),
		system.NewField("_dependence", fieldtype.VARCHAR, 32, "", false, false, false, "Dependencia", nil),
		system.NewField("_zone", fieldtype.VARCHAR, 32, "", false, false, false, "Zona", nil),
		system.NewField("_phone", fieldtype.TEL, -1, "", false, false, false, "Tel&eacute;fono", nil),

		system.NewField("_caipad", fieldtype.INT, -1, "0", false, false, false, "Matrícula", nil),
	})

	// geo:enrollments7
	var cGeoEnrollments7 = geo.NewGenericController("Enrollments7", true)

	cGeoEnrollments7.NewGenericModel(action.GET, []*system.Field{
		system.NewField("_id", fieldtype.INT, -1, "0", false, true, true, "ID", nil),
		system.NewField("_regional", fieldtype.VARCHAR, 32, "", false, false, false, "Regional", nil),
		system.NewField("_circuit", fieldtype.VARCHAR, 2, "00", false, false, false, "Circuito", nil),
		system.NewField("_budgetcode", fieldtype.VARCHAR, 4, "", true, false, false, "C&oacute;digo", nil),
		system.NewField("_name", fieldtype.VARCHAR, 64, "0000", false, false, false, "Instituci&oacute;n", nil),
		system.NewField("_town", fieldtype.VARCHAR, 256, "", false, false, false, "Poblado", nil),
		system.NewField("_dependence", fieldtype.VARCHAR, 32, "", false, false, false, "Dependencia", nil),
		system.NewField("_zone", fieldtype.VARCHAR, 32, "", false, false, false, "Zona", nil),
		system.NewField("_phone", fieldtype.TEL, -1, "", false, false, false, "Tel&eacute;fono", nil),

		system.NewField("_en", fieldtype.INT, -1, "0", false, false, false, "Matrícula", nil),
	})

	// geo:enrollments8
	var cGeoEnrollments8 = geo.NewGenericController("Enrollments8", true)

	cGeoEnrollments8.NewGenericModel(action.GET, []*system.Field{
		system.NewField("_id", fieldtype.INT, -1, "-1", false, true, true, "ID", nil),
		system.NewField("_regional", fieldtype.VARCHAR, 32, "", false, false, false, "Regional", nil),
		system.NewField("_circuit", fieldtype.VARCHAR, 2, "00", false, false, false, "Circuito", nil),
		system.NewField("_budgetcode", fieldtype.VARCHAR, 4, "", true, false, false, "C&oacute;digo", nil),
		system.NewField("_name", fieldtype.VARCHAR, 64, "0000", false, false, false, "Instituci&oacute;n", nil),
		system.NewField("_town", fieldtype.VARCHAR, 256, "", false, false, false, "Poblado", nil),
		system.NewField("_dependence", fieldtype.VARCHAR, 32, "", false, false, false, "Dependencia", nil),
		system.NewField("_zone", fieldtype.VARCHAR, 32, "", false, false, false, "Zona", nil),
		system.NewField("_phone", fieldtype.TEL, -1, "", false, false, false, "Tel&eacute;fono", nil),

		system.NewField("_cl", fieldtype.INT, -1, "0", false, false, false, "CL", nil),
		system.NewField("_ec", fieldtype.INT, -1, "0", false, false, false, "EC", nil),
		system.NewField("_pn", fieldtype.INT, -1, "0", false, false, false, "PN", nil),
	})

	// geo:enrollments9
	var cGeoEnrollments9 = geo.NewGenericController("Enrollments9", true)

	cGeoEnrollments9.NewGenericModel(action.GET, []*system.Field{
		system.NewField("_id", fieldtype.INT, -1, "0", false, true, true, "ID", nil),
		system.NewField("_regional", fieldtype.VARCHAR, 32, "", false, false, false, "Regional", nil),
		system.NewField("_circuit", fieldtype.VARCHAR, 2, "00", false, false, false, "Circuito", nil),
		system.NewField("_budgetcode", fieldtype.VARCHAR, 4, "", true, false, false, "C&oacute;digo", nil),
		system.NewField("_name", fieldtype.VARCHAR, 64, "0000", false, false, false, "Instituci&oacute;n", nil),
		system.NewField("_town", fieldtype.VARCHAR, 256, "", false, false, false, "Poblado", nil),
		system.NewField("_dependence", fieldtype.VARCHAR, 32, "", false, false, false, "Dependencia", nil),
		system.NewField("_zone", fieldtype.VARCHAR, 32, "", false, false, false, "Zona", nil),
		system.NewField("_phone", fieldtype.TEL, -1, "", false, false, false, "Tel&eacute;fono", nil),

		system.NewField("_ec", fieldtype.INT, -1, "0", false, false, false, "EC", nil),
		system.NewField("_ee", fieldtype.INT, -1, "0", false, false, false, "EE", nil),
		system.NewField("_pea", fieldtype.INT, -1, "0", false, false, false, "PEA", nil),
	})

	// geo:enrollments10
	var cGeoEnrollments10 = geo.NewGenericController("Enrollments10", true)

	cGeoEnrollments10.NewGenericModel(action.GET, []*system.Field{
		system.NewField("_id", fieldtype.INT, -1, "-1", false, true, true, "ID", nil),
		system.NewField("_regional", fieldtype.VARCHAR, 32, "", false, false, false, "Regional", nil),
		system.NewField("_circuit", fieldtype.VARCHAR, 2, "00", false, false, false, "Circuito", nil),
		system.NewField("_budgetcode", fieldtype.VARCHAR, 4, "", true, false, false, "C&oacute;digo", nil),
		system.NewField("_name", fieldtype.VARCHAR, 64, "0000", false, false, false, "Instituci&oacute;n", nil),
		system.NewField("_town", fieldtype.VARCHAR, 256, "", false, false, false, "Poblado", nil),
		system.NewField("_dependence", fieldtype.VARCHAR, 32, "", false, false, false, "Dependencia", nil),
		system.NewField("_zone", fieldtype.VARCHAR, 32, "", false, false, false, "Zona", nil),
		system.NewField("_phone", fieldtype.TEL, -1, "", false, false, false, "Tel&eacute;fono", nil),

		system.NewField("_coned", fieldtype.INT, -1, "0", false, false, false, "Matrícula", nil),
	})
	/* -------------------------------------------------------------------------------------------------------------------------------- */
	// catalog:categories
	var catalog = merp.NewModule("Catalog", "Merp", "../system")

	var cCatalogCategories = catalog.NewGenericController("Categories", false)
	var mCatalogCategories_id = system.NewField("_id", fieldtype.INT, -1, "", false, true, true, "Identificador", nil)
	var mCatalogCategories_title = system.NewField("_title", fieldtype.VARCHAR, 64, "", false, false, false, "T&iacute;tulo", nil)
	var mCatalogCategories = cCatalogCategories.NewGenericModel(action.GET, []*system.Field{
		mCatalogCategories_id,
		mCatalogCategories_title,
		system.NewField("_category", fieldtype.VARCHAR, 64, "", true, false, false, "S&uacute;per-categor&iacute;a", nil),
	})

	var eCatalogCategories = system.NewEnum(mCatalogCategories, mCatalogCategories_id, mCatalogCategories_title)
	var fCatalogCategories = []*system.Field{
		mCatalogCategories_id,
		mCatalogCategories_title,
		system.NewField("_category", fieldtype.INT, -1, "0", true, false, false, "S&uacute;per-categor&iacute;a", eCatalogCategories),
	}

	cCatalogCategories.NewGenericModel(action.SET, fCatalogCategories)

	// catalog:products
	var cCatalogProducts = catalog.NewGenericController("Products", false)
	var mCatalogProducts_id = system.NewField("_id", fieldtype.INT, -1, "", false, true, true, "Identificador", nil)
	var mCatalogProducts_title = system.NewField("_title", fieldtype.VARCHAR, 64, "", false, false, false, "T&iacute;tulo", nil)
	var mCatalogProducts_href = system.NewField("_href", fieldtype.VARCHAR, 1024, "", false, false, false, "Enlace", nil)

	cCatalogProducts.NewGenericModel(action.GET, []*system.Field{
		mCatalogProducts_id,
		mCatalogProducts_title,
		mCatalogProducts_href,
		system.NewField("_category", fieldtype.VARCHAR, 64, "", true, false, false, "Categor&iacute;a", nil),
	})
	cCatalogProducts.NewGenericModel(action.SET, []*system.Field{
		mCatalogProducts_id,
		mCatalogProducts_title,
		mCatalogProducts_href,
		system.NewField("_category", fieldtype.INT, -1, "", true, false, false, "Categor&iacute;a", eCatalogCategories),
	})
	/* -------------------------------------------------------------------------------------------------------------------------------- */
	var matrix = merp.NewModule("Matrix", "Merp", "../system")

	// matrix:indicators
	var cMatrixIndicators = matrix.NewGenericController("Indicators", false)
	cMatrixIndicators.NewGenericModel(action.SET, []*system.Field{
		system.NewField("_id", fieldtype.INT, -1, "", false, true, true, "Identificador", nil),
		system.NewField("_description", fieldtype.VARCHAR, 512, "", false, false, false, "Descripci&oacute;n", nil),
	})
	/* -------------------------------------------------------------------------------------------------------------------------------- */
	var info = merp.NewModule("Info", "Merp", "../system")

	// info:requests
	var cInfoRequests = info.NewGenericController("Requests", false)
	var mInfoRequests_id = system.NewField("_id", fieldtype.INT, -1, "", false, true, true, "Identificador", nil)
	cInfoRequests.NewGenericModel(action.GET, []*system.Field{
		mInfoRequests_id,
		system.NewField("_applicant", fieldtype.TEXT, -1, "", false, false, false, "Identificaci&oacute;n del solicitante", nil),
		system.NewField("_description", fieldtype.TEXT, -1, "", false, false, false, "Informaci&oacute;n de la solicitud", nil),
		system.NewField("_detail", fieldtype.TEXT, -1, "", false, false, false, "Informaci&oacute;n relevante", nil),
		system.NewField("_mailto", fieldtype.VARCHAR, 128, "", false, false, false, "Receptor", nil),
		system.NewField("_mailsent", fieldtype.VARCHAR, 4, "", false, false, false, "Enviado", nil),
		system.NewField("_recvackn", fieldtype.TEXT, -1, "", true, false, false, "Acuse de recibido", nil),
	})

	cInfoRequests.NewGenericModel(action.SET, []*system.Field{
		mInfoRequests_id,
		system.NewField("_recvdate", fieldtype.DATE, -1, "", false, false, false, "Fecha de recepci&oacute;n", nil),
		system.NewField("_recvunit", fieldtype.VARCHAR, 64, "", false, false, false, "Unidad receptora", nil),
		system.NewField("_recvreceptor", fieldtype.VARCHAR, 128, "", false, false, false, "Nombre del funcionario", nil),
		system.NewField("_recvnotified", fieldtype.VARCHAR, 128, "", false, false, false, "Notificado a", nil),
	})
	/* -------------------------------------------------------------------------------------------------------------------------------- */
	var security = merp.NewModule("Security", "Merp", "../system")

	// security:profiles
	var cSecurityProfiles = security.NewGenericController("Profiles", false)
	var mSecurityProfiles_id = system.NewField("_id", fieldtype.INT, -1, "", false, true, true, "Identificador", nil)
	var mSecurityProfiles_name = system.NewField("_name", fieldtype.VARCHAR, 64, "", false, false, false, "Nombre", nil)
	cSecurityProfiles.NewGenericModel(action.GET, []*system.Field{
		mSecurityProfiles_id,
		mSecurityProfiles_name,
		system.NewField("_reserved", fieldtype.VARCHAR, 4, "", false, false, false, "Reservado", nil),
	})

	cSecurityProfiles.NewGenericModel(action.SET, []*system.Field{
		mSecurityProfiles_id,
		mSecurityProfiles_name,
		// system.NewField("_reserved", fieldtype.BIT, -1, "0", false, false, false, "Reservado para el personal", nil),
	})

	// security:users
	var cSecurityUsers = security.NewGenericController("Users", false)
	var mSecurityUsers_id = system.NewField("_id", fieldtype.INT, -1, "", false, true, true, "Identificador", nil)
	var mSecurityUsers_email = system.NewField("_email", fieldtype.VARCHAR, 32, "", false, false, false, "e-Mail", nil)
	var mSecurityUsers_phone = system.NewField("_phone", fieldtype.INT, -1, "", true, false, false, "Tel&eacute;fono", nil)
	var mSecurityUsers_photo = system.NewField("_photo", fieldtype.VARCHAR, 32, "", true, false, false, "Foto", nil)
	cSecurityUsers.NewGenericModel(action.GET, []*system.Field{
		mSecurityUsers_id,
		system.NewField("_fullname", fieldtype.VARCHAR, 512, "", false, false, false, "Nombre completo", nil),
		mSecurityUsers_email,
		mSecurityUsers_phone,
		mSecurityUsers_photo,
		system.NewField("_status", fieldtype.VARCHAR, 4, "", false, false, false, "Estado", nil),
	})

	cSecurityUsers.NewGenericModel(action.SET, []*system.Field{
		mSecurityUsers_id,
		system.NewField("_1stname", fieldtype.VARCHAR, 32, "", false, false, false, "Nombre", nil),
		system.NewField("_2ndname", fieldtype.VARCHAR, 32, "", true, false, false, "Segundo nombre", nil),
		system.NewField("_3rdname", fieldtype.VARCHAR, 32, "", false, false, false, "Apellido", nil),
		system.NewField("_4thname", fieldtype.VARCHAR, 32, "", true, false, false, "Segundo apellido", nil),
		mSecurityUsers_email,
		mSecurityUsers_phone,
		mSecurityUsers_photo,
		// system.NewField("_status", fieldtype.INT, -1, "", false, false, false, "Estado", nil),
	})
	/* -------------------------------------------------------------------------------------------------------------------------------- */
	var sys = merp.NewModule("Sys", "Merp", "../system")

	// sys:settings
	var cSysSettings = sys.NewGenericController("Settings", false)
	cSysSettings.NewGenericModel(action.SET, []*system.Field{
		system.NewField("_id", fieldtype.INT, -1, "", false, true, true, "Identificador", nil),
		system.NewField("_key", fieldtype.VARCHAR, 32, "", false, false, false, "Llave", nil),
		system.NewField("_value", fieldtype.VARCHAR, 64, "", false, false, false, "Valor", nil),
	})

	merp.Publish("C:\\xampp\\htdocs\\phactory")
}
