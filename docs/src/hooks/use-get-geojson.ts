import { validateAndFixGeoJson } from '@/utils/fotmat-geojson';
import { useQuery } from '@tanstack/react-query';
import axios from 'axios';
import type { MultiPolygon } from 'geojson';

interface GeoJsonLayerProps {
  geojsonUrl?: string;
  enabled: boolean;
}

const RETRY_WHEN_FAILED = 1;
const CACHE_TIME = 5 * 60 * 1000;

const getGeoJSON = async (url: string) => {
  const response = await axios.get<unknown>(url);
  const fixedData = validateAndFixGeoJson(response.data);
  return fixedData;
};

export const useFetchGeoJSON = ({
  geojsonUrl,
  enabled = true,
}: GeoJsonLayerProps) => {
  return useQuery<MultiPolygon | null>({
    queryKey: ['geojson', geojsonUrl],
    queryFn: () => getGeoJSON(geojsonUrl ?? ''),
    enabled: enabled && !!geojsonUrl,
    staleTime: CACHE_TIME,
    retry: RETRY_WHEN_FAILED,
  });
};
