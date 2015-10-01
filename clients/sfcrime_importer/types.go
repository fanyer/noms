// This file was generated by nomdl/codegen.

package main

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __mainPackageInFile_types_CachedRef = __mainPackageInFile_types_Ref()

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func __mainPackageInFile_types_Ref() ref.Ref {
	p := types.PackageDef{
		NamedTypes: types.MapOfStringToTypeRefDef{

			"Geoposition": __typeRefForGeoposition,
			"Incident":    __typeRefForIncident,
		},
	}.New()
	return types.RegisterPackage(&p)
}

// ListOfIncident

type ListOfIncident struct {
	l types.List
}

func NewListOfIncident() ListOfIncident {
	return ListOfIncident{types.NewList()}
}

type ListOfIncidentDef []IncidentDef

func (def ListOfIncidentDef) New() ListOfIncident {
	l := make([]types.Value, len(def))
	for i, d := range def {
		l[i] = d.New().NomsValue()
	}
	return ListOfIncident{types.NewList(l...)}
}

func (l ListOfIncident) Def() ListOfIncidentDef {
	d := make([]IncidentDef, l.Len())
	for i := uint64(0); i < l.Len(); i++ {
		d[i] = IncidentFromVal(l.l.Get(i)).Def()
	}
	return d
}

func ListOfIncidentFromVal(val types.Value) ListOfIncident {
	// TODO: Validate here
	return ListOfIncident{val.(types.List)}
}

func (l ListOfIncident) NomsValue() types.Value {
	return l.l
}

func (l ListOfIncident) Equals(p ListOfIncident) bool {
	return l.l.Equals(p.l)
}

func (l ListOfIncident) Ref() ref.Ref {
	return l.l.Ref()
}

// A Noms Value that describes ListOfIncident.
var __typeRefForListOfIncident = types.MakeCompoundTypeRef("", types.ListKind, types.MakeTypeRef("Incident", ref.Ref{}))

func (m ListOfIncident) TypeRef() types.TypeRef {
	return __typeRefForListOfIncident
}

func (l ListOfIncident) Len() uint64 {
	return l.l.Len()
}

func (l ListOfIncident) Empty() bool {
	return l.Len() == uint64(0)
}

func (l ListOfIncident) Get(i uint64) Incident {
	return IncidentFromVal(l.l.Get(i))
}

func (l ListOfIncident) Slice(idx uint64, end uint64) ListOfIncident {
	return ListOfIncident{l.l.Slice(idx, end)}
}

func (l ListOfIncident) Set(i uint64, val Incident) ListOfIncident {
	return ListOfIncident{l.l.Set(i, val.NomsValue())}
}

func (l ListOfIncident) Append(v ...Incident) ListOfIncident {
	return ListOfIncident{l.l.Append(l.fromElemSlice(v)...)}
}

func (l ListOfIncident) Insert(idx uint64, v ...Incident) ListOfIncident {
	return ListOfIncident{l.l.Insert(idx, l.fromElemSlice(v)...)}
}

func (l ListOfIncident) Remove(idx uint64, end uint64) ListOfIncident {
	return ListOfIncident{l.l.Remove(idx, end)}
}

func (l ListOfIncident) RemoveAt(idx uint64) ListOfIncident {
	return ListOfIncident{(l.l.RemoveAt(idx))}
}

func (l ListOfIncident) fromElemSlice(p []Incident) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v.NomsValue()
	}
	return r
}

type ListOfIncidentIterCallback func(v Incident, i uint64) (stop bool)

func (l ListOfIncident) Iter(cb ListOfIncidentIterCallback) {
	l.l.Iter(func(v types.Value, i uint64) bool {
		return cb(IncidentFromVal(v), i)
	})
}

type ListOfIncidentIterAllCallback func(v Incident, i uint64)

func (l ListOfIncident) IterAll(cb ListOfIncidentIterAllCallback) {
	l.l.IterAll(func(v types.Value, i uint64) {
		cb(IncidentFromVal(v), i)
	})
}

type ListOfIncidentFilterCallback func(v Incident, i uint64) (keep bool)

func (l ListOfIncident) Filter(cb ListOfIncidentFilterCallback) ListOfIncident {
	nl := NewListOfIncident()
	l.IterAll(func(v Incident, i uint64) {
		if cb(v, i) {
			nl = nl.Append(v)
		}
	})
	return nl
}

// Incident

type Incident struct {
	m types.Map
}

func NewIncident() Incident {
	return Incident{types.NewMap(
		types.NewString("$name"), types.NewString("Incident"),
		types.NewString("$type"), types.MakeTypeRef("Incident", __mainPackageInFile_types_CachedRef),
		types.NewString("ID"), types.Int64(0),
		types.NewString("Category"), types.NewString(""),
		types.NewString("Description"), types.NewString(""),
		types.NewString("DayOfWeek"), types.NewString(""),
		types.NewString("Date"), types.NewString(""),
		types.NewString("Time"), types.NewString(""),
		types.NewString("PdDistrict"), types.NewString(""),
		types.NewString("Resolution"), types.NewString(""),
		types.NewString("Address"), types.NewString(""),
		types.NewString("Geoposition"), NewGeoposition().NomsValue(),
		types.NewString("PdID"), types.NewString(""),
	)}
}

type IncidentDef struct {
	ID          int64
	Category    string
	Description string
	DayOfWeek   string
	Date        string
	Time        string
	PdDistrict  string
	Resolution  string
	Address     string
	Geoposition GeopositionDef
	PdID        string
}

func (def IncidentDef) New() Incident {
	return Incident{
		types.NewMap(
			types.NewString("$name"), types.NewString("Incident"),
			types.NewString("$type"), types.MakeTypeRef("Incident", __mainPackageInFile_types_CachedRef),
			types.NewString("ID"), types.Int64(def.ID),
			types.NewString("Category"), types.NewString(def.Category),
			types.NewString("Description"), types.NewString(def.Description),
			types.NewString("DayOfWeek"), types.NewString(def.DayOfWeek),
			types.NewString("Date"), types.NewString(def.Date),
			types.NewString("Time"), types.NewString(def.Time),
			types.NewString("PdDistrict"), types.NewString(def.PdDistrict),
			types.NewString("Resolution"), types.NewString(def.Resolution),
			types.NewString("Address"), types.NewString(def.Address),
			types.NewString("Geoposition"), def.Geoposition.New().NomsValue(),
			types.NewString("PdID"), types.NewString(def.PdID),
		)}
}

func (s Incident) Def() (d IncidentDef) {
	d.ID = int64(s.m.Get(types.NewString("ID")).(types.Int64))
	d.Category = s.m.Get(types.NewString("Category")).(types.String).String()
	d.Description = s.m.Get(types.NewString("Description")).(types.String).String()
	d.DayOfWeek = s.m.Get(types.NewString("DayOfWeek")).(types.String).String()
	d.Date = s.m.Get(types.NewString("Date")).(types.String).String()
	d.Time = s.m.Get(types.NewString("Time")).(types.String).String()
	d.PdDistrict = s.m.Get(types.NewString("PdDistrict")).(types.String).String()
	d.Resolution = s.m.Get(types.NewString("Resolution")).(types.String).String()
	d.Address = s.m.Get(types.NewString("Address")).(types.String).String()
	d.Geoposition = GeopositionFromVal(s.m.Get(types.NewString("Geoposition"))).Def()
	d.PdID = s.m.Get(types.NewString("PdID")).(types.String).String()
	return
}

// A Noms Value that describes Incident.
var __typeRefForIncident = types.MakeStructTypeRef("Incident",
	[]types.Field{
		types.Field{"ID", types.MakePrimitiveTypeRef(types.Int64Kind), false},
		types.Field{"Category", types.MakePrimitiveTypeRef(types.StringKind), false},
		types.Field{"Description", types.MakePrimitiveTypeRef(types.StringKind), false},
		types.Field{"DayOfWeek", types.MakePrimitiveTypeRef(types.StringKind), false},
		types.Field{"Date", types.MakePrimitiveTypeRef(types.StringKind), false},
		types.Field{"Time", types.MakePrimitiveTypeRef(types.StringKind), false},
		types.Field{"PdDistrict", types.MakePrimitiveTypeRef(types.StringKind), false},
		types.Field{"Resolution", types.MakePrimitiveTypeRef(types.StringKind), false},
		types.Field{"Address", types.MakePrimitiveTypeRef(types.StringKind), false},
		types.Field{"Geoposition", types.MakeTypeRef("Geoposition", ref.Ref{}), false},
		types.Field{"PdID", types.MakePrimitiveTypeRef(types.StringKind), false},
	},
	types.Choices{},
)

func (m Incident) TypeRef() types.TypeRef {
	return __typeRefForIncident
}

func IncidentFromVal(val types.Value) Incident {
	// TODO: Validate here
	return Incident{val.(types.Map)}
}

func (s Incident) NomsValue() types.Value {
	return s.m
}

func (s Incident) Equals(other Incident) bool {
	return s.m.Equals(other.m)
}

func (s Incident) Ref() ref.Ref {
	return s.m.Ref()
}

func (s Incident) Type() types.TypeRef {
	return s.m.Get(types.NewString("$type")).(types.TypeRef)
}

func (s Incident) ID() int64 {
	return int64(s.m.Get(types.NewString("ID")).(types.Int64))
}

func (s Incident) SetID(val int64) Incident {
	return Incident{s.m.Set(types.NewString("ID"), types.Int64(val))}
}

func (s Incident) Category() string {
	return s.m.Get(types.NewString("Category")).(types.String).String()
}

func (s Incident) SetCategory(val string) Incident {
	return Incident{s.m.Set(types.NewString("Category"), types.NewString(val))}
}

func (s Incident) Description() string {
	return s.m.Get(types.NewString("Description")).(types.String).String()
}

func (s Incident) SetDescription(val string) Incident {
	return Incident{s.m.Set(types.NewString("Description"), types.NewString(val))}
}

func (s Incident) DayOfWeek() string {
	return s.m.Get(types.NewString("DayOfWeek")).(types.String).String()
}

func (s Incident) SetDayOfWeek(val string) Incident {
	return Incident{s.m.Set(types.NewString("DayOfWeek"), types.NewString(val))}
}

func (s Incident) Date() string {
	return s.m.Get(types.NewString("Date")).(types.String).String()
}

func (s Incident) SetDate(val string) Incident {
	return Incident{s.m.Set(types.NewString("Date"), types.NewString(val))}
}

func (s Incident) Time() string {
	return s.m.Get(types.NewString("Time")).(types.String).String()
}

func (s Incident) SetTime(val string) Incident {
	return Incident{s.m.Set(types.NewString("Time"), types.NewString(val))}
}

func (s Incident) PdDistrict() string {
	return s.m.Get(types.NewString("PdDistrict")).(types.String).String()
}

func (s Incident) SetPdDistrict(val string) Incident {
	return Incident{s.m.Set(types.NewString("PdDistrict"), types.NewString(val))}
}

func (s Incident) Resolution() string {
	return s.m.Get(types.NewString("Resolution")).(types.String).String()
}

func (s Incident) SetResolution(val string) Incident {
	return Incident{s.m.Set(types.NewString("Resolution"), types.NewString(val))}
}

func (s Incident) Address() string {
	return s.m.Get(types.NewString("Address")).(types.String).String()
}

func (s Incident) SetAddress(val string) Incident {
	return Incident{s.m.Set(types.NewString("Address"), types.NewString(val))}
}

func (s Incident) Geoposition() Geoposition {
	return GeopositionFromVal(s.m.Get(types.NewString("Geoposition")))
}

func (s Incident) SetGeoposition(val Geoposition) Incident {
	return Incident{s.m.Set(types.NewString("Geoposition"), val.NomsValue())}
}

func (s Incident) PdID() string {
	return s.m.Get(types.NewString("PdID")).(types.String).String()
}

func (s Incident) SetPdID(val string) Incident {
	return Incident{s.m.Set(types.NewString("PdID"), types.NewString(val))}
}

// Geoposition

type Geoposition struct {
	m types.Map
}

func NewGeoposition() Geoposition {
	return Geoposition{types.NewMap(
		types.NewString("$name"), types.NewString("Geoposition"),
		types.NewString("$type"), types.MakeTypeRef("Geoposition", __mainPackageInFile_types_CachedRef),
		types.NewString("Latitude"), types.Float32(0),
		types.NewString("Longitude"), types.Float32(0),
	)}
}

type GeopositionDef struct {
	Latitude  float32
	Longitude float32
}

func (def GeopositionDef) New() Geoposition {
	return Geoposition{
		types.NewMap(
			types.NewString("$name"), types.NewString("Geoposition"),
			types.NewString("$type"), types.MakeTypeRef("Geoposition", __mainPackageInFile_types_CachedRef),
			types.NewString("Latitude"), types.Float32(def.Latitude),
			types.NewString("Longitude"), types.Float32(def.Longitude),
		)}
}

func (s Geoposition) Def() (d GeopositionDef) {
	d.Latitude = float32(s.m.Get(types.NewString("Latitude")).(types.Float32))
	d.Longitude = float32(s.m.Get(types.NewString("Longitude")).(types.Float32))
	return
}

// A Noms Value that describes Geoposition.
var __typeRefForGeoposition = types.MakeStructTypeRef("Geoposition",
	[]types.Field{
		types.Field{"Latitude", types.MakePrimitiveTypeRef(types.Float32Kind), false},
		types.Field{"Longitude", types.MakePrimitiveTypeRef(types.Float32Kind), false},
	},
	types.Choices{},
)

func (m Geoposition) TypeRef() types.TypeRef {
	return __typeRefForGeoposition
}

func GeopositionFromVal(val types.Value) Geoposition {
	// TODO: Validate here
	return Geoposition{val.(types.Map)}
}

func (s Geoposition) NomsValue() types.Value {
	return s.m
}

func (s Geoposition) Equals(other Geoposition) bool {
	return s.m.Equals(other.m)
}

func (s Geoposition) Ref() ref.Ref {
	return s.m.Ref()
}

func (s Geoposition) Type() types.TypeRef {
	return s.m.Get(types.NewString("$type")).(types.TypeRef)
}

func (s Geoposition) Latitude() float32 {
	return float32(s.m.Get(types.NewString("Latitude")).(types.Float32))
}

func (s Geoposition) SetLatitude(val float32) Geoposition {
	return Geoposition{s.m.Set(types.NewString("Latitude"), types.Float32(val))}
}

func (s Geoposition) Longitude() float32 {
	return float32(s.m.Get(types.NewString("Longitude")).(types.Float32))
}

func (s Geoposition) SetLongitude(val float32) Geoposition {
	return Geoposition{s.m.Set(types.NewString("Longitude"), types.Float32(val))}
}
