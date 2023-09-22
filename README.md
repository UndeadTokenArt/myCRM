myCRM is a small program that runs locally using the Gin framework for GO.


features:
- A front page with a CSS style Sheet
- NewCustomerForm page for adding a client to the Database
- 


Technologies
- Go programing language
- SQlite - Database
- Gin framework for handing router
- GORM - managing Data base connection



ToDo:
- Auto populating templates and forms with Client data
- User account managment
    - store User contact info
- templated messages
    - word replacement
    - Send email
- Calander
- add import/export of CSV files

Installation instructions
1. install GO - https://webinstall.dev/golang/
    Windows - $ curl.exe https://webi.ms/golang | powershell
    Mac - $ curl -sS https://webi.sh/golang | sh
    Linux - $ curl -sS https://webi.sh/golang | sh
2. clone down the repo and all dependencies
    $ git clone https://github.com/UndeadTokenArt/myCRM.git
    $ go get 
3. in the terminal line
    $ go run .