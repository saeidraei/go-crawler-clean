# Url Scrapper in golang using clean architecture
###how to run using docker-compose:
all you need to do is run `docker-compose up` in project root directory and after some time that images are pulled and some of them are build , project is up and running.

you can use the [postman docs](https://documenter.getpostman.com/view/3010056/SzKTvyo2?version=latest) to explore the api.

### How failed jobs are being handled?
if request to a url fails the failedCount attribute of url would be incremented and the url would be added at end of queue if the failedCount is not more than 3.

if the url content had no title the no title attribute would be set to true.

### Clean Architecture :
Layers ( from the most abstract to the most concrete ) :
- domain : abstract data structures
- uc : "use cases", the pure business logic
- implem : implementations of the interfaces used in the business logic (uc layer)
- infra : setup/configuration of the implementation
