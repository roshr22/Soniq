func handleUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // Limit = 10MB

	file, handler, err := r.FormFile("audio")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Save the file locally
	dst, err := os.Create("uploads/" + handler.Filename)
	if err != nil {
		http.Error(w, "Unable to save the file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	io.Copy(dst, file)

	// Respond with file URL
	fileURL := "/uploads/" + handler.Filename
	w.Write([]byte(fileURL))
}

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/ws", handleWebSocket)
	router.HandleFunc("/upload", handleUpload).Methods("POST")

	// Serve uploaded files
	router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads/"))))
}
