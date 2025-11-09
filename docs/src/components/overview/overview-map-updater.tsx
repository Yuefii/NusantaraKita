import { useEffect } from 'react';
import { useMap } from 'react-leaflet';

interface MapUpdaterProps {
  center: [number, number];
  zoom: number;
}

const MapUpdater = ({ center, zoom }: MapUpdaterProps) => {
  const map = useMap();
  useEffect(() => {
    map.setView(center, zoom);
  }, [center, zoom, map]);
  return null;
};

export default MapUpdater;
