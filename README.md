# batch csv to http post

This is sort of a one day project I made when I had an issue where a ui element accidentally let somebody create a whole bunch of things in an application, and there was no way in the app to undo that action, which was terribly inconvenient for the user, so I built this tool to  batch update all the items which had been created.

This tool works well when you have direct access to the database but don't want to just change data due to possible side effects.

your csv file might look something like:

```
client_id,firstname,middlename,lastname
3456,Michael,Wayne,Wilding
3555,Jane,W,Austen

```

The endpoint to update the client info might look like `http://hostname/update-client`

and the form data might look like `client_id=###&firstname=string&middlename=string&lastname=string`

where changing any of the values will update the record in the database.

the tool will run through every line of the csv document
and make a post with the data on each row

so the example from the csv above would look like

`/update-client?client_id=3465&firstname=Michael&middlename=Wayne&lastname=Wilding`

# Config
- build and install by running `go install`
- if your environment is setup, you will now be able to run `csvToHttpPost`
- if you run without an argument you will get a readme including an example schema for your config.
- copy the schema to an actual json file and fill in the info.  Note it includes an array of headers.  You can add as many headers as you want.  I used it to add the appropriate session cookies so I can authenticate my calls.  Also, there is a known bug where you cant add more than one cookie.  My mvp didnt need it so I choose to save that edge case for the future. 
- If you run `csvToHttpPost /path/to/file` you will see requests and responses printed to stdout.
---
# Alternate Config
For convenience I have added a binary compiled for mac.  Now you can run it as long as you have the right permissions and put in a directory accessible from your environment PATH.

