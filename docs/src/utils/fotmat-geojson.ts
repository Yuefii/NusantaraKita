import type { MultiPolygon, Polygon } from 'geojson';

type ValidatedGeoJson = {
  type: 'MultiPolygon';
  coordinates: unknown;
};

function validateGeoJsonInput(data: unknown): ValidatedGeoJson {
  if (typeof data !== 'object' || data === null) {
    throw new Error('Data GeoJSON bukan sebuah objek');
  }

  if (!('type' in data && 'coordinates' in data)) {
    throw new Error("Objek GeoJSON tidak memiliki 'type' atau 'coordinates'");
  }

  const geoData = data as { type: unknown; coordinates: unknown };

  if (geoData.type !== 'MultiPolygon') {
    throw new Error(
      `Tipe GeoJSON yang diharapkan 'MultiPolygon', tetapi menerima '${geoData.type}'`,
    );
  }

  return geoData as ValidatedGeoJson;
}

// function to check if coordinates is a polygon structure
function isPolygonStructure(
  coordinates: unknown,
): coordinates is Polygon['coordinates'] {
  try {
    const c = coordinates as unknown[][][];
    return (
      Array.isArray(c) &&
      Array.isArray(c[0]) &&
      Array.isArray(c[0][0]) &&
      typeof c[0][0][0] === 'number'
    );
  } catch {
    return false;
  }
}

// function to check if coordinates is a multipolygon structure
function isValidMultiPolygonStructure(
  coordinates: unknown,
): coordinates is MultiPolygon['coordinates'] {
  try {
    const c = coordinates as unknown[][][][];
    return (
      Array.isArray(c) &&
      Array.isArray(c[0]) &&
      Array.isArray(c[0][0]) &&
      Array.isArray(c[0][0][0]) &&
      typeof c[0][0][0][0] === 'number'
    );
  } catch {
    return false;
  }
}

function validateAndFixGeoJson(data: unknown): MultiPolygon {
  const geoData = validateGeoJsonInput(data);

  // convert polygon to multipolygon
  if (isPolygonStructure(geoData.coordinates)) {
    return {
      type: 'MultiPolygon',
      coordinates: [geoData.coordinates], // process convert polygon to multipolygon
    };
  }

  // if already multipolygon return as is
  if (isValidMultiPolygonStructure(geoData.coordinates)) {
    return data as MultiPolygon;
  }

  throw new Error(
    "Data memiliki tipe 'MultiPolygon' tetapi struktur koordinat tidak dikenal",
  );
}

export { validateAndFixGeoJson };
