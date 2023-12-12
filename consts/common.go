package consts

import "time"

// Collections 结构体表示集合的结构。
type Collections struct {
	// Collections 字段表示集合的列表。
	Collections []struct {
		// Id 字段表示集合的唯一标识。
		Id string `json:"id"`
		// Title 字段表示集合的标题。
		Title string `json:"title"`
		// CanRead 字段表示是否可以读取集合。
		CanRead bool `json:"can_read"`
		// CanWrite 字段表示是否可以写入集合。
		CanWrite bool `json:"can_write"`
		// MediaTypes 字段表示集合支持的媒体类型。
		MediaTypes []string `json:"media_types"`
	} `json:"collections"`
}

// ApiRootsRes 结构体保存了API根路径的相关信息
type ApiRootsRes struct {
	Title    string   `json:"title"`     // 标题字段用于表示API根路径的标题
	Default  string   `json:"default"`   // 默认字段用于表示API根路径的默认值
	ApiRoots []string `json:"api_roots"` // api_roots字段用于保存API根路径的集合
}

// GetCollectionDataRes 结构体用于表示获取集合数据的响应结果
type GetCollectionDataRes struct {
	// more 表示是否还有更多数据
	More bool `json:"more"`
	// objects 是一个数组，用于存储具体的数据对象
	Objects []struct {
		// id 是数据对象的唯一标识符
		Id string `json:"id"`
		// type 是数据对象的类型
		Type string `json:"type"`
		// spec_version 是数据对象的规范版本
		SpecVersion string `json:"spec_version"`
		// created 是数据对象创建的时间
		Created time.Time `json:"created"`
		// modified 是数据对象最后修改的时间
		Modified time.Time `json:"modified"`
		// name 是数据对象的名称（可选字段）
		Name string `json:"name,omitempty"`
		// malware_types 是该数据对象关联的恶意软件类型（可选字段）
		MalwareTypes []string `json:"malware_types,omitempty"`
		// is_family 表示该数据对象是否是一个家族（可选字段）
		IsFamily bool `json:"is_family,omitempty"`
		// relationship_type 是该数据对象的关系类型（可选字段）
		RelationshipType string `json:"relationship_type,omitempty"`
		// source_ref 是该数据对象的源引用（可选字段）
		SourceRef string `json:"source_ref,omitempty"`
		// target_ref 是该数据对象的目标引用（可选字段）
		TargetRef string `json:"target_ref,omitempty"`
		// description 是该数据对象的描述（可选字段）
		Description string `json:"description,omitempty"`
		// pattern 是该数据对象的模式（可选字段）
		Pattern string `json:"pattern,omitempty"`
		// pattern_type 是该数据对象的模式类型（可选字段）
		PatternType string `json:"pattern_type,omitempty"`
		// pattern_version 是该数据对象的模式版本（可选字段）
		PatternVersion string `json:"pattern_version,omitempty"`
		// valid_from 是该数据对象有效的起始时间（可选字段）
		ValidFrom time.Time `json:"valid_from,omitempty"`
		// extensions 是该数据对象的扩展属性（可选字段）
		Extensions map[string]interface{} `json:"extensions,omitempty"`
		// ioc 是该数据对象的指示性特征（可选字段）
		Ioc []Ioc `json:"ioc,omitempty"`
		// ioc_category 是该数据对象的指示性特征类别（可选字段）
		IocCategory string `json:"ioc_category,omitempty"`
		// targeted 是该数据对象的目标（可选字段）
		Targeted interface{} `json:"targeted,omitempty"`
		// risk 是该数据对象的风险等级（可选字段）
		Risk string `json:"risk,omitempty"`
		// malicious_type 是该数据对象的恶意类型（可选字段）
		MaliciousType string `json:"malicious_type,omitempty"`
		// judge 是该数据对象的判断结果（可选字段）
		Judge string `json:"judge,omitempty"`
		// ttp 是该数据对象的战术、技巧和过程（可选字段）
		Ttp string `json:"ttp,omitempty"`
		// campaign 是该数据对象关联的campaign（可选字段）
		Campaign []string `json:"campaign,omitempty"`
		// malicious_family 是该数据对象关联的恶意家族（可选字段）
		MaliciousFamily []string `json:"malicious_family,omitempty"`
		// labels 是该数据对象的标签（可选字段）
		Labels []string `json:"labels,omitempty"`
		// industry_sector 是该数据对象所属的行业部门（可选字段）
		IndustrySector []string `json:"industry_sector,omitempty"`
		// platform 是该数据对象支持的平台（可选字段）
		Platform []string `json:"platform,omitempty"`
		// created_by_ref 是该数据对象创建者的引用（可选字段）
		CreatedByRef string `json:"created_by_ref,omitempty"`
		// schema 是该数据对象的架构（可选字段）
		Schema string `json:"schema,omitempty"`
		// version 是该数据对象的版本号（可选字段）
		Version string `json:"version,omitempty"`
		// extension_types 是该数据对象的扩展类型（可选字段）
		ExtensionTypes []string `json:"extension_types,omitempty"`
		// extension_properties 是该数据对象的扩展属性（可选字段）
		ExtensionProperties []string `json:"extension_properties,omitempty"`
	} `json:"objects"`
	// next 是下一个数据集合的链接
	Next string `json:"next"`
}

// Ioc 结构体表示指示性特征对象
type Ioc struct {
	// 值字段表示指示性特征的值
	Value string `json:"value,omitempty"`
	// 类型字段表示指示性特征的类型
	Stype string `json:"stype,omitempty"`
}
