
package virtualmachineconfig

import (
    "github.com/hashicorp/terraform/helper/schema"
    nutanixV3 "nutanixV3"
    "time"
    "strings"
)

func convertToBool(a interface{}) bool {
    if a != nil {
        return a.(bool)
    }
    return false
}

func convertToInt(a interface{}) int64 {
    if a != nil {
        i :=  a.(int)
        return int64(i)
    }
    return 0
}

func convertToString(a interface{}) string {
    if a != nil {
        return a.(string)
    }
    return ""
}

// SetMachineConfig function sets fields in struct from ResourceData
func SetMachineConfig(d *schema.ResourceData) nutanixV3.VmIntentInput {
    spec := d.Get("spec").(*schema.Set).List()           // spec
    metadata := d.Get("metadata").(*schema.Set).List()                     // metadata
    machine := nutanixV3.VmIntentInput{
        ApiVersion: "3.0", // api_version
        Spec:       SetSpec(spec, 0),   //Spec
        Metadata:   SetMetadata(metadata, 0),   //Metadata
    }
    if strings.ToUpper(machine.Spec.Resources.PowerState) == "ON" {
        machine.Spec.Resources.PowerState = "POWERED_ON"
    } else {
        machine.Spec.Resources.PowerState = "POWERED_OFF"
    }
    machine.Metadata.Kind = "vm"
    machine.Spec.Name = d.Get("name").(string)
    machine.Metadata.Name = d.Get("name").(string)
    return machine
}



// SetClusterReference sets cluster_reference fields in json struct
func SetClusterReference (t []interface{}, i int) nutanixV3.ClusterReference {
	if len(t) > 0 {
		s := t[i].(map[string]interface{})

		 

		ClusterReference := nutanixV3.ClusterReference{
			Name:		convertToString(s["name"]),
			Uuid:		convertToString(s["uuid"]),
			Kind:		convertToString(s["kind"]),

		}
		return ClusterReference
	}
	return nutanixV3.ClusterReference{}
}


// SetGpuList sets gpu_list fields in json struct
func SetGpuList (t []interface{}, i int) nutanixV3.VmGpu {
	if len(t) > 0 {
		s := t[i].(map[string]interface{})

		 

		GpuList := nutanixV3.VmGpu{
			Mode:		convertToString(s["mode"]),
			DeviceId:		convertToInt(s["device_id"]),
			Vendor:		convertToString(s["vendor"]),

		}
		return GpuList
	}
	return nutanixV3.VmGpu{}
}


// SetSysprep sets sysprep fields in json struct
func SetSysprep (t []interface{}, i int) nutanixV3.GuestCustomizationSysprep {
	if len(t) > 0 {
		s := t[i].(map[string]interface{})

		 

		Sysprep := nutanixV3.GuestCustomizationSysprep{
			InstallType:		convertToString(s["install_type"]),
			UnattendXml:		convertToString(s["unattend_xml"]),

		}
		return Sysprep
	}
	return nutanixV3.GuestCustomizationSysprep{}
}


// SetCloudInit sets cloud_init fields in json struct
func SetCloudInit (t []interface{}, i int) nutanixV3.GuestCustomizationCloudInit {
	if len(t) > 0 {
		s := t[i].(map[string]interface{})

		 

		CloudInit := nutanixV3.GuestCustomizationCloudInit{
			MetaData:		convertToString(s["meta_data"]),
			UserData:		convertToString(s["user_data"]),

		}
		return CloudInit
	}
	return nutanixV3.GuestCustomizationCloudInit{}
}


// SetGuestCustomization sets guest_customization fields in json struct
func SetGuestCustomization (t []interface{}, i int) nutanixV3.GuestCustomization {
	if len(t) > 0 {
		s := t[i].(map[string]interface{})

		 

		GuestCustomization := nutanixV3.GuestCustomization{
			Sysprep:		SetSysprep(s["sysprep"].(*schema.Set).List(), 0),
			CloudInit:		SetCloudInit(s["cloud_init"].(*schema.Set).List(), 0),

		}
		return GuestCustomization
	}
	return nutanixV3.GuestCustomization{}
}


// SetDiskAddress sets disk_address fields in json struct
func SetDiskAddress (t []interface{}, i int) nutanixV3.DiskAddress {
	if len(t) > 0 {
		s := t[i].(map[string]interface{})

		 

		DiskAddress := nutanixV3.DiskAddress{
			AdapterType:		convertToString(s["adapter_type"]),
			DeviceIndex:		convertToInt(s["device_index"]),

		}
		return DiskAddress
	}
	return nutanixV3.DiskAddress{}
}


// SetBootConfig sets boot_config fields in json struct
func SetBootConfig (t []interface{}, i int) nutanixV3.VmBootConfig {
	if len(t) > 0 {
		s := t[i].(map[string]interface{})

		 

		BootConfig := nutanixV3.VmBootConfig{
			DiskAddress:		SetDiskAddress(s["disk_address"].(*schema.Set).List(), 0),
			MacAddress:		convertToString(s["mac_address"]),

		}
		return BootConfig
	}
	return nutanixV3.VmBootConfig{}
}


// SetDataSourceReference sets data_source_reference fields in json struct
func SetDataSourceReference (t []interface{}, i int) nutanixV3.Reference {
	if len(t) > 0 {
		s := t[i].(map[string]interface{})

		 

		DataSourceReference := nutanixV3.Reference{
			Kind:		convertToString(s["kind"]),
			Uuid:		convertToString(s["uuid"]),
			Name:		convertToString(s["name"]),

		}
		return DataSourceReference
	}
	return nutanixV3.Reference{}
}


// SetDeviceProperties sets device_properties fields in json struct
func SetDeviceProperties (t []interface{}, i int) nutanixV3.VmDiskDeviceProperties {
	if len(t) > 0 {
		s := t[i].(map[string]interface{})

		 

		DeviceProperties := nutanixV3.VmDiskDeviceProperties{
			DiskAddress:		SetDiskAddress(s["disk_address"].(*schema.Set).List(), 0),
			DeviceType:		convertToString(s["device_type"]),

		}
		return DeviceProperties
	}
	return nutanixV3.VmDiskDeviceProperties{}
}


// SetDiskList sets disk_list fields in json struct
func SetDiskList (t []interface{}, i int) nutanixV3.VmDisk {
	if len(t) > 0 {
		s := t[i].(map[string]interface{})

		 

		DiskList := nutanixV3.VmDisk{
			DiskSizeMib:		convertToInt(s["disk_size_mib"]),
			DataSourceReference:		SetDataSourceReference(s["data_source_reference"].(*schema.Set).List(), 0),
			DeviceProperties:		SetDeviceProperties(s["device_properties"].(*schema.Set).List(), 0),

		}
		return DiskList
	}
	return nutanixV3.VmDisk{}
}


// SetIPEndpointList sets ip_endpoint_list fields in json struct
func SetIPEndpointList (t []interface{}, i int) nutanixV3.IpAddress {
	if len(t) > 0 {
		s := t[i].(map[string]interface{})

		 

		IPEndpointList := nutanixV3.IpAddress{
			Type:		convertToString(s["type"]),
			Address:		convertToString(s["address"]),

		}
		return IPEndpointList
	}
	return nutanixV3.IpAddress{}
}


// SetNetworkFunctionChainReference sets network_function_chain_reference fields in json struct
func SetNetworkFunctionChainReference (t []interface{}, i int) nutanixV3.NetworkFunctionChainReference {
	if len(t) > 0 {
		s := t[i].(map[string]interface{})

		 

		NetworkFunctionChainReference := nutanixV3.NetworkFunctionChainReference{
			Kind:		convertToString(s["kind"]),
			Name:		convertToString(s["name"]),
			Uuid:		convertToString(s["uuid"]),

		}
		return NetworkFunctionChainReference
	}
	return nutanixV3.NetworkFunctionChainReference{}
}


// SetSubnetReference sets subnet_reference fields in json struct
func SetSubnetReference (t []interface{}, i int) nutanixV3.SubnetReference {
	if len(t) > 0 {
		s := t[i].(map[string]interface{})

		 

		SubnetReference := nutanixV3.SubnetReference{
			Kind:		convertToString(s["kind"]),
			Name:		convertToString(s["name"]),
			Uuid:		convertToString(s["uuid"]),

		}
		return SubnetReference
	}
	return nutanixV3.SubnetReference{}
}


// SetNicList sets nic_list fields in json struct
func SetNicList (t []interface{}, i int) nutanixV3.VmNic {
	if len(t) > 0 {
		s := t[i].(map[string]interface{})

		

		var  IPEndpointList []nutanixV3.IpAddress
		if s["ip_endpoint_list"] != nil {
			for i := 0; i< len(s["ip_endpoint_list"].([]interface{})); i++ {
				elem := SetIPEndpointList(s["ip_endpoint_list"].([]interface{}), i)
				IPEndpointList = append(IPEndpointList, elem)
			}
		}

 

		NicList := nutanixV3.VmNic{
			NetworkFunctionNicType:		convertToString(s["network_function_nic_type"]),
			MacAddress:		convertToString(s["mac_address"]),
			IpEndpointList:		IPEndpointList,
			NetworkFunctionChainReference:		SetNetworkFunctionChainReference(s["network_function_chain_reference"].(*schema.Set).List(), 0),
			NicType:		convertToString(s["nic_type"]),
			SubnetReference:		SetSubnetReference(s["subnet_reference"].(*schema.Set).List(), 0),

		}
		return NicList
	}
	return nutanixV3.VmNic{}
}


// SetParentReference sets parent_reference fields in json struct
func SetParentReference (t []interface{}, i int) nutanixV3.Reference {
	if len(t) > 0 {
		s := t[i].(map[string]interface{})

		 

		ParentReference := nutanixV3.Reference{
			Kind:		convertToString(s["kind"]),
			Uuid:		convertToString(s["uuid"]),
			Name:		convertToString(s["name"]),

		}
		return ParentReference
	}
	return nutanixV3.Reference{}
}


// SetResources sets resources fields in json struct
func SetResources (t []interface{}, i int) nutanixV3.VmResources {
	if len(t) > 0 {
		s := t[i].(map[string]interface{})

		

		var  GpuList []nutanixV3.VmGpu
		if s["gpu_list"] != nil {
			for i := 0; i< len(s["gpu_list"].([]interface{})); i++ {
				elem := SetGpuList(s["gpu_list"].([]interface{}), i)
				GpuList = append(GpuList, elem)
			}
		}



		var  DiskList []nutanixV3.VmDisk
		if s["disk_list"] != nil {
			for i := 0; i< len(s["disk_list"].([]interface{})); i++ {
				elem := SetDiskList(s["disk_list"].([]interface{}), i)
				DiskList = append(DiskList, elem)
			}
		}



		var  NicList []nutanixV3.VmNic
		if s["nic_list"] != nil {
			for i := 0; i< len(s["nic_list"].([]interface{})); i++ {
				elem := SetNicList(s["nic_list"].([]interface{}), i)
				NicList = append(NicList, elem)
			}
		}

 

		Resources := nutanixV3.VmResources{
			NumVcpusPerSocket:		convertToInt(s["num_vcpus_per_socket"]),
			MemorySizeMb:		convertToInt(s["memory_size_mb"]),
			GpuList:		GpuList,
			GuestCustomization:		SetGuestCustomization(s["guest_customization"].(*schema.Set).List(), 0),
			BootConfig:		SetBootConfig(s["boot_config"].(*schema.Set).List(), 0),
			DiskList:		DiskList,
			NicList:		NicList,
			PowerState:		convertToString(s["power_state"]),
			NumSockets:		convertToInt(s["num_sockets"]),
			ParentReference:		SetParentReference(s["parent_reference"].(*schema.Set).List(), 0),

		}
		return Resources
	}
	return nutanixV3.VmResources{}
}


// SetSpec sets spec fields in json struct
func SetSpec (t []interface{}, i int) nutanixV3.Vm {
	if len(t) > 0 {
		s := t[i].(map[string]interface{})

		 

		Spec := nutanixV3.Vm{
			Description:		convertToString(s["description"]),
			ClusterReference:		SetClusterReference(s["cluster_reference"].(*schema.Set).List(), 0),
			Resources:		SetResources(s["resources"].(*schema.Set).List(), 0),
			Name:		convertToString(s["name"]),

		}
		return Spec
	}
	return nutanixV3.Vm{}
}


// SetCategories sets categories fields in json struct
func SetCategories(s map[string]interface{}) map[string]string {
	var CategoriesI map[string]interface{}
	if s["categories"] != nil {
		CategoriesI = s["categories"].(map[string]interface{})
	}
	Categories := make(map[string]string)
	for key, value := range CategoriesI {
		switch value := value.(type){
		case string:
			Categories[key]	= value
		}
	}
	return Categories
}


// SetOwnerReference sets owner_reference fields in json struct
func SetOwnerReference (t []interface{}, i int) nutanixV3.UserReference {
	if len(t) > 0 {
		s := t[i].(map[string]interface{})

		 

		OwnerReference := nutanixV3.UserReference{
			Uuid:		convertToString(s["uuid"]),
			Kind:		convertToString(s["kind"]),
			Name:		convertToString(s["name"]),

		}
		return OwnerReference
	}
	return nutanixV3.UserReference{}
}


// SetMetadata sets metadata fields in json struct
func SetMetadata (t []interface{}, i int) nutanixV3.VmMetadata {
	if len(t) > 0 {
		s := t[i].(map[string]interface{})

		

		var CreationTime time.Time
		tempCreationTime := convertToString(s["creation_time"])
		if tempCreationTime != ""{
			CreationTime, _ = time.Parse(tempCreationTime, tempCreationTime)
		}



		var LastUpdateTime time.Time
		tempLastUpdateTime := convertToString(s["last_update_time"])
		if tempLastUpdateTime != ""{
			LastUpdateTime, _ = time.Parse(tempLastUpdateTime, tempLastUpdateTime)
		}

 

		Metadata := nutanixV3.VmMetadata{
			Kind:		convertToString(s["kind"]),
			Uuid:		convertToString(s["uuid"]),
			CreationTime:		CreationTime,
			Categories:		SetCategories(s),
			OwnerReference:		SetOwnerReference(s["owner_reference"].(*schema.Set).List(), 0),
			EntityVersion:		convertToInt(s["entity_version"]),
			Name:		convertToString(s["name"]),
			LastUpdateTime:		LastUpdateTime,

		}
		return Metadata
	}
	return nutanixV3.VmMetadata{}
}


// SetVMIntentInput sets VmIntentInput fields in json struct
func SetVMIntentInput (t []interface{}, i int) nutanixV3.VmIntentInput {
	if len(t) > 0 {
		s := t[i].(map[string]interface{})

		 

		VMIntentInput := nutanixV3.VmIntentInput{
			Spec:		SetSpec(s["spec"].(*schema.Set).List(), 0),
			ApiVersion:		convertToString(s["api_version"]),
			Metadata:		SetMetadata(s["metadata"].(*schema.Set).List(), 0),

		}
		return VMIntentInput
	}
	return nutanixV3.VmIntentInput{}
}
