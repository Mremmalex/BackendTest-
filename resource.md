# reference for the program

[golang sql1](https://flaviocopes.com/golang-sql-database/)
[how to limit a database query](https://gist.github.com/ssokolow/262503)

## TODOs

### Public Area

1. __A public event list:__ When accessing the main route, the application should show a list of all the events registered, paging them every 10 events;

- 1.1 The user should be able to filter the list of events by dates, or regions;

2. __Event details:__ the application must allow the user to see the details of the event, by clicking on the event listing, or accessing the event link;

3. __signup:__ the application should allow the user to register by informing: Name, Email, Password, Bio, Profile Picture, City, and State;

4. __User login:__ The application should allow the user to login using their credentials;

- 4.1 The login should persist when the application is closed, or reloaded;

### Logged Area

5. __Friend invitation:__ the application will allow the user to enter an email to add as a friend;

6. __Add as friend:__ The informed user should receive a friend request, or an invitation to register, if they are not already a user;

7. __Friendship management:__ the user will be able to see your new friend requests, list your friends, and undo friendships;

8. __Event registration:__ the application should allow the user to register an event by informing: Name, Description, Date, Time, and Place;

- 8.1 The user should be able to edit and cancel events their events;

9. __Invite friends to event:__ the user can invite their friends to events, being able to invite all friends, or only the selected ones;

- 9.1 If the user has already been invited to the event, regardless of their status (confirmed, rejected, awaiting confirmation), the invited user should not be notified of the invitation again;

10. __My event list:__ the user should be able to see their events, being able to filter them by those who will participate, and the ones that he created;

11. __Manage event invitations:__ The user can accept, or reject, attend events.

12. __Events management:__ The user can view their rejected events and undo rejections, deciding to participate, if the event has not yet occurred;

