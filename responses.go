package yandexDiskApiClient

import (
	"fmt"
	"time"
)

type SuccessResponse struct {
	OperationId string `json:"operation_Id"`
	Href        string `json:"href"`
	Method      string `json:"method"`
	Templated   bool   `json:"templated"`
}

func (e *SuccessResponse) Info() string {
	return fmt.Sprintf("OperationId: %s\nHref: %s\nMethod: %s\nTemplated: %s\n", e.OperationId, e.Href, e.Method, e.Templated)
}

type ErrorResponse struct {
	Message     string `json:"message"`
	Description string `json:"description"`
	Error       string `json:"error"`
}

func (e *ErrorResponse) Info() string {
	return fmt.Sprintf("message: %s\ndescription: %s\nerror: %s\n", e.Message, e.Description, e.Error)
}

type Disk struct {
	MaxFileSize                int           `json:"max_file_size"`
	PaidMaxFileSize            int64         `json:"paid_max_file_size"`
	TotalSpace                 int64         `json:"total_space"`
	TrashSize                  int           `json:"trash_size"`
	IsPaid                     bool          `json:"is_paid"`
	UsedSpace                  int64         `json:"used_space"`
	SystemFolders              SystemFolders `json:"system_folders"`
	User                       User          `json:"user"`
	UnlimitedAutouploadEnabled bool          `json:"unlimited_autoupload_enabled"`
	Revision                   int64         `json:"revision"`
}

func (d *Disk) Info() string {
	return fmt.Sprintf("UserName: %s\nUsedSpace: %d\nUsedSpace: %d\n", d.User.DisplayName, d.UsedSpace, d.TotalSpace)
}

type SystemFolders struct {
	Odnoklassniki string `json:"odnoklassniki"`
	Google        string `json:"google"`
	Instagram     string `json:"instagram"`
	Vkontakte     string `json:"vkontakte"`
	Mailru        string `json:"mailru"`
	Downloads     string `json:"downloads"`
	Applications  string `json:"applications"`
	Facebook      string `json:"facebook"`
	Social        string `json:"social"`
	Scans         string `json:"scans"`
	Screenshots   string `json:"screenshots"`
	Photostream   string `json:"photostream"`
}

type User struct {
	Country     string `json:"country"`
	Login       string `json:"login"`
	DisplayName string `json:"display_name"`
	UID         string `json:"uid"`
}

type FilesResourceList struct {
	Items  []Resource `json:"items"`
	Limit  int        `json:"limit"`
	Offset int        `json:"offset"`
}

func (f *FilesResourceList) Info() string {
	result := ""
	for _, file := range f.Items {
		result += file.Path + "\n"
	}
	result += fmt.Sprintf("files count: %d", len(f.Items))

	return result
}

type TrashResourceList struct {
	Embedded struct {
		Sort   string     `json:"sort"`
		Items  []Resource `json:"items"`
		Limit  int        `json:"limit"`
		Offset int        `json:"offset"`
		Path   string     `json:"path"`
		Total  int        `json:"total"`
	} `json:"_embedded"`
}

func (t *TrashResourceList) Info() string {
	result := ""
	for _, file := range t.Embedded.Items {
		result += t.Embedded.Path + file.Name + "\n"
	}
	result += fmt.Sprintf("files count: %d", t.Embedded.Total)

	return result
}

type Resource struct {
	AntivirusStatus string     `json:"antivirus_status"`
	Size            int        `json:"size"`
	CommentIds      CommentIds `json:"comment_ids"`
	Name            string     `json:"name"`
	Exif            Exif       `json:"exif,omitempty"`
	Created         time.Time  `json:"created"`
	ResourceID      string     `json:"resource_id"`
	Modified        time.Time  `json:"modified"`
	MimeType        string     `json:"mime_type"`
	File            string     `json:"file"`
	MediaType       string     `json:"media_type"`
	Preview         string     `json:"preview"`
	Path            string     `json:"path"`
	Sha256          string     `json:"sha256"`
	Type            string     `json:"type"`
	Md5             string     `json:"md5"`
	Revision        int64      `json:"revision"`
	OriginPath      string     `json:"origin_path"`
	Deleted         time.Time  `json:"deleted"`
	PublicUrl       string     `json:"public_Url"`
	PublicKey       string     `json:"public_Key"`
}

type CommentIds struct {
	PrivateResource string `json:"private_resource"`
	PublicResource  string `json:"public_resource"`
}

type Exif struct {
	DateTime time.Time `json:"date_time"`
}
