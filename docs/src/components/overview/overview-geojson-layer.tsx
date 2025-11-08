import { useFetchGeoJSON } from '@/hooks/use-get-geojson';
import { GeoJSON } from 'react-leaflet';

interface GeoJsonLayerProps {
  geojsonUrl?: string;
  colorHex?: string;
  weight?: number;
  fillOpacity?: number;
}

const GeoJsonLayer = ({
  geojsonUrl,
  colorHex,
  fillOpacity,
  weight,
}: GeoJsonLayerProps) => {
  const {
    data: geoData,
    isError,
    isLoading,
  } = useFetchGeoJSON({
    geojsonUrl,
    enabled: !!geojsonUrl,
  });

  if (isLoading) return null;
  if (isError || !geoData) return null;

  return (
    <GeoJSON
      data={geoData}
      style={() => ({
        color: colorHex || '#3388ff',
        weight: weight ?? 2,
        fillOpacity: fillOpacity ?? 0.2,
      })}
    />
  );
};

export default GeoJsonLayer;
