package servicenow

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Err represents a possible error message that came back from the server
type Err struct {
	Err    string `json:"error"`
	Reason string `json:"reason"`
}

func (e Err) Error() string {
	if e.Reason == "" {
		return e.Err
	}
	return fmt.Sprintf("%s: %s", e.Err, e.Reason)
}

// Client helps users interact with a ServiceNow instance.
type Client struct {
	Token, Instance string
}

func sys(param string) string {
	return fmt.Sprintf("sysparm_%s", param)
}

// PerformFor creates and executes an authenticated HTTP request to an instance,
// for the given table, action and optional id, with the passed options, and
// unmarshal's the JSON into the passed output interface pointer, returning an
// error.
func (c *Client) PerformFor(table, action, id string, opts url.Values, body interface{}, out interface{}) error {
	inst := c.Instance

	u := fmt.Sprintf("%s/%s", inst, table)

	meth := http.MethodGet

	req, err := http.NewRequest(meth, u+"?"+opts.Encode(), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", c.Token))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	return json.Unmarshal(data, out)
}

// GetFor performs a servicenow get to the specified table, with options, and
// unmarshal's JSON into the output parameter.
func (c Client) GetFor(table string, id string, opts url.Values, out interface{}) error {
	return c.PerformFor(table, "get", id, opts, nil, out)
}

// GetRecordsFor performs a servicenow getRecords to the specified table, with
// options, and unmarshal's JSON into the output parameter.
func (c Client) GetRecordsFor(table string, opts url.Values, out interface{}) error {
	return c.PerformFor(table, "getRecords", "", opts, nil, out)
}

// GetRecords performs a servicenow getRecords to the specified table, with
// options, and returns a map for each item
func (c Client) GetRecords(table string, opts url.Values) ([]map[string]interface{}, error) {
	var out struct {
		Records []map[string]interface{}
	}
	err := c.GetRecordsFor(table, opts, &out)
	return out.Records, err
}

// Insert creates a new record for the specified table, with the specified obj
// data, and takes a destination object out for the response data.
func (c Client) Insert(table string, obj, out interface{}) error {
	return c.PerformFor(table, "insert", "", nil, obj, out)
}

// Update creates a new record for the specified table, with the specified obj
// data, and takes a destination object out for the response data.
func (c Client) Update(table string, opts url.Values, obj, out interface{}) error {
	return c.PerformFor(table, "update", "", opts, obj, out)
}
