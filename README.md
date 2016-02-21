# DigiEase
These are the services I made for an application called DigiEase which helps citizens of Pune to fill in various police forms.

The module DigiEaseCitizen are a collection of services which will be used by the application on citizen's end 
and DigiEasePolice are a collection of services which will be used by the application on police's end.

The services for only one form are included right now - the tenant registration form.

The modules are both MVC based the database used to store data is MySQL.

The control flow of the application is as follows:

  1. The request arrives on a specific path. The router routes the request to the an appropriate controller.
  2. The controller breaks the request and gets the input parameters. It the passes the parameters to respective model function
     which calculates and maintains data.
  3. The processed data is sent back to controller which formats it in JSON and sends back on the route.
  
If you want to see the data and the processing part, direclty jump to models.

You can contact me anytime for any queries
