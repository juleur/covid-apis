export default interface DepartmentDelimit {
    type: "Point" | "MultiPoint" | "LineString" | "MultiLineString" | "Polygon" | "MultiPolygon" | "GeometryCollection" | "Feature" | "FeatureCollection";
    features: Feature[];
}

interface Feature {
    type: string;
    geometry: Geometry;
    properties: Properties;
}

interface Geometry {
    type: string;
    coordinates: number[][][];
}

interface Properties {
    code: string;
    nom: string;
}
