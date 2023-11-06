# Earthquake Notifier

The feed is been provided by the geophysics university of greece and its an RSS feed with the latest data of earthquakes in the Greek Space. It only displays the last 112 earthquakes.

The Notifier takes two types of parameters before deciding to send an alert:

#### The Filters:
```json
{
  "parameters": [
    {
      "MaxDistanseInKM": 64,
      "Timestamp": "",
      "MinDepth": 3,
      "MinMagnitude": 3.6
    },
    {
      "MaxDistanseInKM": 85,
      "Timestamp": "",
      "MinDepth": 3,
      "MinMagnitude": 1
    }
  ]
}

```

 - MaxDistanseInKM: Its the max distanse from your location that you want to be notified about the earthquake based on the rest of the options
 - Timestamp: is not implemented yet
 - MinDepth: Minimum depth of the earthquake
 - MinMagnitude: Minimum Magnitude to be notified about

#### Your location in Longitude and Latitude
At the moment it is statically configued (Work In Progress)

## TODO:
- [ ] Use Channels to send data to SendAlerts func
- [ ] Use Go Routines to Send Notifications
- [ ] Dynamic Location
- [ ] Check for Duplicates
- [ ] Implement Push Notifications
- [ ] Provide a UI
