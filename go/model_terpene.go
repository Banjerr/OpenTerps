/*
 * OpenTerps
 *
 * An Open API that contains information about terpenes, the effects, and the cannabis varieties that contain them.
 *
 * API version: 0.0.1
 * Contact: benjamminredden@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Terpene struct {
	Id int64 `json:"id,omitempty"`
	Category *Category `json:"category,omitempty"`
	Name string `json:"name"`
	Tags []Tag `json:"tags,omitempty"`
	Tastes []Taste `json:"tastes,omitempty"`
	Smells []Smell `json:"smells,omitempty"`
	Strains []Strain `json:"strains,omitempty"`
}