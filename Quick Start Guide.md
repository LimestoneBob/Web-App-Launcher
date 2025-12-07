This app lets you draw or load a boundary on a map, pull building outlines from OpenStreetMap, and review area, height, and other details in a table you can export to CSV or GeoJSON.[1]

***

## 1. Opening the app

- Open the `PropertyAssetReview.htm` file in a modern browser such as Chrome or Edge.[1]  
- Make the browser window fairly large so you can see both the map and the data panel at the bottom.[1]

***

## 2. Basic layout and controls

- The main screen shows a map, a search bar across the top, and a slim data panel along the bottom that can be expanded.[1]  
- A small menu button on the left opens a side panel with tools for drawing a search area and clearing the map.[1]

***

## 3. Choosing country and region

- At the top, small chips let you switch between Canada and the United States and cycle through key provinces or states such as ON, QC, BC or CA, TX, NY.[1]  
- Changing the country will also move the map view to a sensible starting city for that country, while region changes adjust filters used in address search and data retrieval.[1]

***

## 4. Finding an area by address

- In the search box, type an address or place name (for example, a street address or building name) and press Enter or click the magnifying glass icon.[1]  
- When a match is found, the map zooms to that location and drops a marker; press the search button again to retrieve nearby building data for that spot.[1]

***

## 5. Drawing your own boundary

- Open the left sidebar and choose “Draw Area Box or Polygon” to turn on the drawing tool, then drag a rectangle or click around a polygon to outline your area of interest.[1]  
- When you finish, the app clears any previous boundary, zooms to your new shape, and starts retrieving buildings that fall mostly inside that boundary.[1]

***

## 6. Loading a boundary from a file

- Click the plus button in the top search bar to open the “Load Data Options” window, choose a GeoJSON file, and then click “Load Boundary.”[1]  
- The boundary from the file is added to the map, and the app automatically starts retrieving and processing buildings within that boundary.[1]

***

## 7. Reviewing the results

- The bottom panel shows total buildings and total area; click the panel header if you want to expand or collapse it.[1]  
- Each row in the table shows a building’s sequence number, name, area, perimeter, height, floor levels, and OpenStreetMap ID; clicking a row highlights that building on the map.[1]

***

## 8. Switching between metric and imperial

- Use the “Metric/Imperial” chip in the panel header to switch units for the table, map pop‑ups, and CSV export.[1]  
- When you switch units, the app refreshes the displayed values so that areas, lengths, and heights use the chosen system.[1]

***

## 9. Exporting your data

- Use the download icons in the panel header to export the current results either as a CSV file (for spreadsheets) or as GeoJSON (for use in mapping tools).[1]  
- CSV exports include each building’s measurements and centroid coordinates, while GeoJSON exports include both footprints and centroid points with useful attributes.[1]

***

## 10. Helpful notice about running a simple local server

- If the app is opened directly from a file (the address bar starts with `file:///`), some online map and data services may not respond because modern browsers block certain requests for security reasons (often called cross‑site restrictions).[2][1]  
- To get the best experience, run the app through a simple local web server so the address bar shows something like `http://localhost:8000` instead of `file:///`; common options include using a lightweight tool in your code editor or a basic command‑line server provided by your operating system.[3][4]  
- If you are connected through a virtual private network, some services that provide maps or location data may still be blocked; if the map tiles or building results do not load as expected, try pausing the virtual private network or switching to a different one and then reload the page.[5][6]

***

[1](https://ppl-ai-file-upload.s3.amazonaws.com/web/direct-files/attachments/20250703/9496716b-e1fb-45c9-992a-7e61a498541c/PropertyAssetReview.htm)  
[2](https://simplelocalize.io/blog/posts/what-is-cors/)  
[3](https://developer.mozilla.org/en-US/docs/Learn_web_development/Howto/Tools_and_setup/set_up_a_local_testing_server)  
[4](https://stackoverflow.com/questions/6084360/using-node-js-as-a-simple-web-server)  
[5](https://fastvpn.com/blog/bypass-geo-blocking/)  
[6](https://www.cape.co/blog/does-a-vpn-hide-your-location)  
[7](https://gist.github.com/jgravois/5e73b56fa7756fd00b89)  
[8](https://learn.microsoft.com/en-us/azure/static-web-apps/get-started-portal)  
[9](https://www.reddit.com/r/webdev/comments/1c6qsib/absolute_beginner_question_but_why_do_we_need_a/)  
[10](https://www.willianantunes.com/blog/2023/07/how-to-test-cors-configuration-locally/)  
[11](https://static-web-server.net)

