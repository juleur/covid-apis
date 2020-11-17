import type CovidReport from '../models/covid-data.model';
import type DepartmentDelimit from '../models/department-delimit.model';

let proxy_on: boolean = true;
let BACKEND_URL: string

if (proxy_on) {
  BACKEND_URL = "http://localhost:8010/proxy/";
} else {
  BACKEND_URL = "https://backend.coro21-jl.xyz/";
}

const fetchConfig: RequestInit = {method: "GET"};

export async function fetchCovidReports(): Promise<{ [key:string]:CovidReport[] }> {
  const response = await fetch(BACKEND_URL+"covid-data", fetchConfig);
  const data: { [key:string]:CovidReport[] }= await response.json();
  return data;
}

export async function fetchGeoJSON(): Promise<DepartmentDelimit> {
  const response = await fetch(BACKEND_URL+"geojson_departements", fetchConfig);
  const data: DepartmentDelimit = await response.json();
  return data; 
}