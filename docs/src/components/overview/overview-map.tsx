import { useOverview, type MapState } from '@/context/overview-provider';
import { cn } from '@/lib/utils';
import { useMemo } from 'react';
import { MapContainer, Marker, Popup, TileLayer } from 'react-leaflet';
import GeoJsonLayer from './overview-geojson-layer';
import MapUpdater from './overview-map-updater';

const DEFAULT_CENTER_POSITION: [number, number] = [-2.5489, 118.0149];

const getActivePosition = (state: MapState): [number, number] => {
  const { desaKel, kecamatan, kabKota, province } = state.selected;
  const levels = [desaKel, kecamatan, kabKota, province];
  const active = levels.find((level) => level && level.lat && level.lng);
  return active ? [active.lat, active.lng] : DEFAULT_CENTER_POSITION;
};

const getZoomLevel = (state: MapState): number => {
  const { desaKel, kecamatan, kabKota, province } = state.selected;
  if (desaKel) return 13;
  if (kecamatan) return 10;
  if (kabKota) return 9;
  if (province) return 8;
  return 6;
};

const OverviewMap = () => {
  const { state } = useOverview();

  const activePosition = useMemo(() => getActivePosition(state), [state]);
  const zoomLevel = useMemo(() => getZoomLevel(state), [state]);

  const activeLabel = useMemo(() => {
    const { desaKel, kecamatan, kabKota, province } = state.selected;
    return (
      desaKel?.nama ||
      kecamatan?.nama ||
      kabKota?.nama ||
      province?.nama ||
      'Indonesia 🇮🇩'
    );
  }, [state.selected]);

  const geojsonLayers = useMemo(() => {
    const { desaKel, kecamatan, kabKota, province } = state.selected;
    return [
      {
        url: province?.geojson_url,
        color: '#1E88E5',
        weight: 2,
        fillOpacity: 0.08,
      },
      {
        url: kabKota?.geojson_url,
        color: '#43A047',
        weight: 2,
        fillOpacity: 0.12,
      },
      {
        url: kecamatan?.geojson_url,
        color: '#FB8C00',
        weight: 2,
        fillOpacity: 0.2,
      },
      {
        url: desaKel?.geojson_url,
        color: '#E53935',
        weight: 2,
        fillOpacity: 0.3,
      },
    ].filter((item) => !!item.url);
  }, [state.selected]);

  return (
    <MapContainer
      center={activePosition}
      zoom={zoomLevel}
      scrollWheelZoom={true}
      className={cn(
        'z-0 aspect-square h-full max-h-[40dvh] w-full rounded-lg md:max-h-[60dvh]',
      )}
    >
      <TileLayer
        attribution='&copy; <a href="https://carto.com/">CartoDB</a>'
        url="https://{s}.basemaps.cartocdn.com/light_all/{z}/{x}/{y}{r}.png"
      />
      <MapUpdater
        center={activePosition}
        zoom={zoomLevel}
      />
      {geojsonLayers.map((layer, index) => (
        <GeoJsonLayer
          key={index}
          geojsonUrl={layer.url}
          colorHex={layer.color}
          weight={layer.weight}
          fillOpacity={layer.fillOpacity}
        />
      ))}
      <Marker position={activePosition}>
        <Popup>{activeLabel}</Popup>
      </Marker>
    </MapContainer>
  );
};

export default OverviewMap;
