package libs

/*
What is the point of the libs module?
1. It will handle interaction between the program and the library files (auth/plan/etc)
2. It will write new files
3. It will read files
4. It will parse the files.

What types of interactions would this module handle?
1. Reading and parsing a file into a struct
    - Decrypt the auth file and library using a local certificate
    - Return a struct with authentication details
2. Storing the parsed data in the auth library so it can be referenced later.
3. Exporting the auth profile to a plan auth file (yaml or something similar)
4. Delete an auth profile from the library
5. Backup the auth library (including or not including a copy of the certificate)
6. Listing the stored auth profiles in the library

- pick a file format (yaml, json, etc)
- pick a library format (yaml, json, sqlite, etc)
- pick an encryption/decryption method (it should be transferable -- meaning the library could be generated on a
    difference system, and then brought over to a new system  as a secret)

Is there a way to do this without having unique fields for each provider?
- could use placeholder field names
  struct Auth {
    fieldA string
    fieldB string
    fieldC string
    ...
 }
- Use
*/
