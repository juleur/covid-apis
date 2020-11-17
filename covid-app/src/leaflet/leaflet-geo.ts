import { geoJSON } from "leaflet";
import Chart from "chart.js";
import type DepartmentDelimit from "../models/department-delimit.model";
import type CovidReport from "../models/covid-data.model";

export function createChart(): HTMLCanvasElement {
  let chartHtml = document.createElement("canvas");
  chartHtml.setAttribute("id", "myChart");
  chartHtml.setAttribute("style", "height:245px;width:325px;");
  return chartHtml;
}

export function buildChart(covidReports: {[key: string]: CovidReport[]}, depName: string): void {
  let canvas = document.getElementById("myChart") as HTMLCanvasElement;
  let ctx = canvas.getContext("2d");
  let data = {
    // récupère les clés du Map
    labels: Object.entries(covidReports).map((i) => i[0]),
    datasets: [
      {
        label: "Guérisons",
        backgroundColor: "rgba(131, 248, 8, .3)",
        borderColor: "rgba(131, 248, 8, 1)",
        hoverBorderColor: "rgba(131, 248, 8, 0.75)",
        data: findCovidDataByDepNameNKey(covidReports, depName, "gueris"),
      },
      {
        label: "Décès",
        backgroundColor: "rgba(67, 67, 67, .3)",
        borderColor: "rgba(67, 67, 67, 1)",
        hoverBorderColor: "rgba(67, 67, 67, 0.75)",
        data: findCovidDataByDepNameNKey(covidReports, depName, "deces"),
      },
      {
        label: "Réanimations",
        backgroundColor: "rgba(177, 13, 253, .3)",
        borderColor: "rgba(177, 13, 253, 1)",
        hoverBorderColor: "rgba(177, 13, 253, 0.75)",
        data: findCovidDataByDepNameNKey(covidReports, depName, "reanimation"),
      },
      {
        label: "Hospitalisations",
        backgroundColor: "rgba(100, 210, 255, .3)",
        borderColor: "rgba(100, 210, 255, 1)",
        hoverBorderColor: "rgba(100, 210, 255, 0.75)",
        data: findCovidDataByDepNameNKey(covidReports, depName, "hospitalises"),
      },
    ],
  };
  let myChart = new Chart(ctx, {
    type: "line",
    data,
  });
  myChart.update();
}

function findCovidDataByDepNameNKey(
    covidReports: {[key: string]: CovidReport[]},
    depName: string,
    key: string
  ): number[] {
    let arr: number[] = [];
    for (const [_, reportByDate] of Object.entries(covidReports)) {
      for (const report of reportByDate) {
        if (report.nom == depName) {
          arr.push(report[key]);
        }
      }
    }
    return arr;
}

export function buildDelimit(geojson: DepartmentDelimit, covidReports: { [key:string]: CovidReport[]}) {
 return geoJSON(geojson, {
      style: function (feature) {
        const lastDatetime = Object.entries(covidReports)[6][0];
        for (let covidReport of covidReports[lastDatetime]) {
            if (covidReport.nom == feature.properties.nom) {
              return {
                opacity: 0.4,
                weight: 3,
                color: "#FFFFFF",
                fillColor: alertLevelColor(covidReport.tauxOccupationReaColor),
                fillOpacity: 0.7,
              };
            }
        }
      },
    })
}

function alertLevelColor(tauxOccup: string): string {
  switch (tauxOccup) {
    case "rouge":
      return "#ff1e1e";
    case "orange":
      return "#ff881e";
    case "vert":
      return "#88ff1e";
    default:
      return "#e5e4e7"
  }
}
