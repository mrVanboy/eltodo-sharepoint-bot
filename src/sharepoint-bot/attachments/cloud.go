package attachments

import (
	"io/ioutil"
	"golang.org/x/oauth2/google"
	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
	"golang.org/x/net/context"
	"encoding/json"
	"sharepoint-bot/cfg"
	"errors"
)

func LoadFromStorage() (Announcements, error) {

	storageData, err := loadJsonFromStorage()
	if err != nil {
		if err == storage.ErrObjectNotExist{
			return nil, nil
		} else {
			return nil, err
		}
	}

	as := Announcements{}

	if err := json.Unmarshal(storageData, &as); err != nil{
		return nil, err
	}
	return as, nil
}

var storageObject *storage.ObjectHandle
var ctx = context.Background()

func loadJsonFromStorage() ([]byte, error) {

	if storageObject == nil {
		err := createStorageObject()
		if err != nil {
			return nil, errors.New(`creating new storage object error: ` + err.Error())
		}
	}


	r, err := storageObject.NewReader(ctx)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func createStorageObject() error{

	credentialsJson := []byte(cfg.Get().GoogleCredentials)
	creds, err := google.CredentialsFromJSON(ctx, credentialsJson, "https://www.googleapis.com/auth/devstorage.read_write")
	if err != nil {
		return err
	}

	client, err := storage.NewClient(ctx, option.WithCredentials(creds))
	if err != nil {
		return err
	}

	bucketName := cfg.Get().GoogleStorageBucket
	bucket := client.Bucket(bucketName)

	storageObject = bucket.Object(cfg.Get().GoogleBucketObject)
	return nil
}

func SaveToStorage(announcements Announcements) error {
	if storageObject == nil {
		err := createStorageObject()
		if err != nil {
			return err
		}
	}

	w := storageObject.NewWriter(ctx)

	jsonWriter := json.NewEncoder(w)
	// jsonWriter := json.NewEncoder(os.Stdout) // for debug only
	err := jsonWriter.Encode(announcements)
	if err != nil {
		return err
	}

	return nil
}