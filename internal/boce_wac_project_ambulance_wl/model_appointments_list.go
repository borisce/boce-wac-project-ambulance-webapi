/*
 * Patients portal for appointments management API
 *
 * Patients portal for appointments management for Web-In-Cloud system
 *
 * API version: 1.0.0
 * Contact: <your_email>
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package boce_wac_project_ambulance_wl

type AppointmentsList struct {

	Id string `json:"id,omitempty"`

	Name string `json:"name"`

	Date string `json:"date"`

	EstimatedStart string `json:"estimatedStart"`

	EstimatedEnd string `json:"estimatedEnd"`

	PatientAppointed bool `json:"patientAppointed"`

	Condition string `json:"condition"`

	DoctorNote string `json:"doctorNote"`
}
