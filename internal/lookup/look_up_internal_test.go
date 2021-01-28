package lookup_test



//func newTestServer(h func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
//	mux := http.NewServeMux()
//	server := httptest.NewServer(mux)
//	//mux.HandleFunc(path, h)
//	return server
//}
//
//func TestInternalNewPhoneLookup(t *testing.T) {
//	uri := "/api/v1/login"
//	server := newTestServer(func(w http.ResponseWriter, r *http.Request) {
//		if (r.URL.String() != "") {
//
//		}
//		//equals(t, req.URL.String(), "/some/path")
//		w.Write([]byte)
//		w.WriteHeader(http.StatusOK)
//	})
//	defer server.Close()
//
//	api := lookup.ClientLookup{
//		Cl: app.InternalAuth{
//			BaseUrl: server.URL,
//			Configuration: struct{ HTTPClient *http.Client }{HTTPClient: server.Client()},
//		},
//	}
//
//
//	//statusCode := lookup.InternalNewPhoneLookup(server.URL, "username", "password")
//	//if statusCode != 200 {
//	//	t.Errorf("Test Login failed")
//	//}
//}