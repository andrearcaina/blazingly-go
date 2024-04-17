# Packages vs Modules vs Libraries

A package is a collection of code (in this context, go files). \
A module is a collection of these packages. \
A library can be considered a collection of these modules or packages that provide specific functionalities towards the end goal of the user. 

An example of a package would be:

```plaintext
random_folder/
├── another_random_folder/
│    └── file_1.go
│    └── file_2.go
└── file_3.go
└── file_4.go               
```

A module would consist of these packages inside a file called `go.mod`.