<script lang="ts">
  import { onMount } from "svelte";
  import { TileLayer, Map, Layer } from "leaflet";
  import { fetchGeoJSON, fetchCovidReports } from "./fetch/fetcher";
  import type DepartmentDelimit from "./models/department-delimit.model";
  import type CovidReport from "./models/covid-data.model";
  import { buildDelimit, createChart, buildChart } from "./leaflet/leaflet-geo";

  let img = document.createElement("img");
  img.setAttribute("style", "width:12px;height:12px");
  img.src = "touchscreen.png";

  onMount(async () => {
    let geoPoint: DepartmentDelimit = await fetchGeoJSON();
    let covidReports: {
      [key: string]: CovidReport[];
    } = await fetchCovidReports();

    let map = new Map("map", {
      center: [46.37582350371957, 2.4758421557273746],
      zoom: 6,
      minZoom: 4,
    });
    const tileLayer = new TileLayer(
      "https://{s}.tile.osm.org/{z}/{x}/{y}.png",
      {
        detectRetina: true,
        attribution:
          '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors',
      }
    );
    tileLayer.addTo(map);

    let geoJSONLayer = buildDelimit(geoPoint, covidReports).addTo(map);
    geoJSONLayer.eachLayer((layer: Layer) => {
      const canvas = createChart();
      layer
        .bindPopup(canvas)
        .bindTooltip(img)
        .on("click", () => {
          let layerDepName: string = layer["feature"]["properties"]["nom"];
          buildChart(covidReports, layerDepName);
        })
        .on("popupclose", () => {
          canvas.remove();
        });
    });
  });
</script>

<style lang="scss">
  .container {
    display: flex;
    flex-wrap: wrap;
    flex-direction: column;
    padding: 2vh 0 2vh 0;
    h3 {
      font-size: 1.5em;
      text-align: center;
      margin-bottom: 3vh;
      text-transform: uppercase;
      color: white;
    }
    #map {
      border-radius: 5px;
      width: 70vw;
      height: 90vh;
    }

    @media only screen and (max-width: 1000px) {
      #map {
        width: 85vw;
        height: 90vh;
      }
    }

    @media only screen and (max-width: 800px) {
      #map {
        width: 96vw;
        height: 90vh;
      }
    }
  }
</style>

<svelte:head>
  <link
    rel="stylesheet"
    href="https://cdn.jsdelivr.net/npm/leaflet@1.7.1/dist/leaflet.css" />
</svelte:head>

<div class="container">
  <h3>Carte Covid19 avec graphique par département</h3>
  <div id="map" />
</div>
