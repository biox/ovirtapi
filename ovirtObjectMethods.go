// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package ovirtapi

import (
	"encoding/xml"
	"reflect"
)

func (api *API) GetVm(id string) (*Vm, error) {
	body, err := api.GetLinkBody(reflect.TypeOf(Vm{}).Name()+"s", id)
	if err != nil {
		return nil, err
	}
	object := api.NewVm()
	err = xml.Unmarshal(body, object)
	if err != nil {
		return nil, err
	}
	return object, err
}

func (api *API) GetAllVm() ([]*Vm, error) {
	body, err := api.GetLinkBody(reflect.TypeOf(Vm{}).Name()+"s", "")
	if err != nil {
		return nil, err
	}
	objects := []*Vm{}
	err = xml.Unmarshal(body, &struct {
		Objects *[]*Vm
	}{&objects})
	if err != nil {
		return nil, err
	}
	for _, object := range objects {
		object.Api = api
	}
	return objects, err
}

func (api *API) NewVm() *Vm {
	return &Vm{OvirtObject: OvirtObject{Api: api}}
}

func (object *Vm) Save() error {
	body, err := xml.MarshalIndent(object, "", "    ")
	if err != nil {
		return err
	}
	// If there is a link, it is an already saved object, we need to update it
	if object.OvirtObject.Href != "" {
		body, err = object.Api.Request("PUT", object.Api.ResolveLink(object.Href), body)
		if err != nil {
			return err
		}
	} else {
		link, err := object.Api.GetLink(reflect.TypeOf(Vm{}).Name() + "s")
		if err != nil {
			return err
		}
		body, err = object.Api.Request("POST", link, body)
		if err != nil {
			return err
		}
	}
	tempObject := Vm{OvirtObject: OvirtObject{Api: object.Api}}
	err = xml.Unmarshal(body, &tempObject)
	*object = tempObject
	if err != nil {
		return err
	}
	return nil
}

func (api *API) GetCluster(id string) (*Cluster, error) {
	body, err := api.GetLinkBody(reflect.TypeOf(Cluster{}).Name()+"s", id)
	if err != nil {
		return nil, err
	}
	object := api.NewCluster()
	err = xml.Unmarshal(body, object)
	if err != nil {
		return nil, err
	}
	return object, err
}

func (api *API) GetAllCluster() ([]*Cluster, error) {
	body, err := api.GetLinkBody(reflect.TypeOf(Cluster{}).Name()+"s", "")
	if err != nil {
		return nil, err
	}
	objects := []*Cluster{}
	err = xml.Unmarshal(body, &struct {
		Objects *[]*Cluster
	}{&objects})
	if err != nil {
		return nil, err
	}
	for _, object := range objects {
		object.Api = api
	}
	return objects, err
}

func (api *API) NewCluster() *Cluster {
	return &Cluster{OvirtObject: OvirtObject{Api: api}}
}

func (object *Cluster) Save() error {
	body, err := xml.MarshalIndent(object, "", "    ")
	if err != nil {
		return err
	}
	// If there is a link, it is an already saved object, we need to update it
	if object.OvirtObject.Href != "" {
		body, err = object.Api.Request("PUT", object.Api.ResolveLink(object.Href), body)
		if err != nil {
			return err
		}
	} else {
		link, err := object.Api.GetLink(reflect.TypeOf(Cluster{}).Name() + "s")
		if err != nil {
			return err
		}
		body, err = object.Api.Request("POST", link, body)
		if err != nil {
			return err
		}
	}
	tempObject := Cluster{OvirtObject: OvirtObject{Api: object.Api}}
	err = xml.Unmarshal(body, &tempObject)
	*object = tempObject
	if err != nil {
		return err
	}
	return nil
}

func (api *API) GetDataCenter(id string) (*DataCenter, error) {
	body, err := api.GetLinkBody(reflect.TypeOf(DataCenter{}).Name()+"s", id)
	if err != nil {
		return nil, err
	}
	object := api.NewDataCenter()
	err = xml.Unmarshal(body, object)
	if err != nil {
		return nil, err
	}
	return object, err
}

func (api *API) GetAllDataCenter() ([]*DataCenter, error) {
	body, err := api.GetLinkBody(reflect.TypeOf(DataCenter{}).Name()+"s", "")
	if err != nil {
		return nil, err
	}
	objects := []*DataCenter{}
	err = xml.Unmarshal(body, &struct {
		Objects *[]*DataCenter
	}{&objects})
	if err != nil {
		return nil, err
	}
	for _, object := range objects {
		object.Api = api
	}
	return objects, err
}

func (api *API) NewDataCenter() *DataCenter {
	return &DataCenter{OvirtObject: OvirtObject{Api: api}}
}

func (object *DataCenter) Save() error {
	body, err := xml.MarshalIndent(object, "", "    ")
	if err != nil {
		return err
	}
	// If there is a link, it is an already saved object, we need to update it
	if object.OvirtObject.Href != "" {
		body, err = object.Api.Request("PUT", object.Api.ResolveLink(object.Href), body)
		if err != nil {
			return err
		}
	} else {
		link, err := object.Api.GetLink(reflect.TypeOf(DataCenter{}).Name() + "s")
		if err != nil {
			return err
		}
		body, err = object.Api.Request("POST", link, body)
		if err != nil {
			return err
		}
	}
	tempObject := DataCenter{OvirtObject: OvirtObject{Api: object.Api}}
	err = xml.Unmarshal(body, &tempObject)
	*object = tempObject
	if err != nil {
		return err
	}
	return nil
}