import { useIsMobile } from '@/hooks/use-mobile';
import { Globe } from 'lucide-react';
import { useMemo } from 'react';
import {
  FaArrowsAltH,
  FaArrowsAltV,
  FaHashtag,
  FaMap,
  FaMapMarkerAlt,
} from 'react-icons/fa';
import { MdErrorOutline } from 'react-icons/md';
import { Alert, AlertDescription, AlertTitle } from '../../ui/alert';
import { CardList } from '../../ui/card-list';
import Code from '../../ui/code';
import { useResponseType } from '../../ui/response-type';
import { CardListSkeleton } from '../../ui/skelekton-card-list';
import SkeletonTable from '../../ui/skelekton-table';
import { Skeleton } from '../../ui/skeleton';
import Table from '../../ui/table';

interface DesaKelProps {
  data: DesaKelApiRes | undefined;
  isError: boolean;
  isLoading: boolean;
  isSuccess: boolean;
  isPending: boolean;
}

const DesaKelTableHeaders = [
  'Kode Kecamatan',
  'Kode Desa/Kel',
  'Nama Desa/Kel',
  'Latitude',
  'Longitude',
  'geojson_url',
];

const parseDesaKelTable = (kabKota: DesaKelApi) => ({
  'Kode Kecamatan': kabKota.kode_kecamatan,
  'Kode Desa/Kel': kabKota.kode,
  'Nama Desa/Kel': kabKota.nama,
  Latitude: kabKota.lat,
  Longitude: kabKota.lng,
  geojson_url: kabKota.geojson_url,
});

const parseDesaKelCard = (kabKota: DesaKelApi) => [
  {
    icon: FaMap,
    title: 'Kode Kecamatan',
    value: kabKota.kode_kecamatan,
  },
  {
    icon: FaHashtag,
    title: 'Kode Desa/Kel',
    value: kabKota.kode,
  },
  {
    icon: FaMapMarkerAlt,
    title: 'Nama Desa/Kel',
    value: kabKota.nama,
  },
  {
    icon: FaArrowsAltV,
    title: 'Latitude',
    value: kabKota.lat.toString(),
  },
  {
    icon: FaArrowsAltH,
    title: 'Longitude',
    value: kabKota.lng.toString(),
  },
  {
    icon: Globe,
    title: 'geojson_url',
    value: kabKota.geojson_url.toString(),
  },
];

const DesaKelErrorAlert = () => {
  return (
    <Alert
      variant="destructive"
      className="flex items-start gap-4 border-red-500 bg-red-100"
    >
      <div className="text-2xl text-red-600">
        <MdErrorOutline />
      </div>
      <div className="flex-1">
        <AlertTitle className="mb-2 text-xl font-bold text-red-500">
          Gagal Memuat Data
        </AlertTitle>
        <AlertDescription className="leading-relaxed text-red-500">
          <span>
            Maaf, terjadi kesalahan saat memuat data.{' '}
            <span
              className="cursor-pointer font-semibold underline"
              onClick={() => window.history.back()}
            >
              Kembali
            </span>
          </span>
        </AlertDescription>
      </div>
    </Alert>
  );
};

export const DesaKelDisplayData = ({
  data,
  isError,
  isLoading,
  isSuccess,
  isPending,
}: DesaKelProps) => {
  const isMobile = useIsMobile();

  const { state } = useResponseType();
  const isShowData = isSuccess && data?.data;

  const parsedDesaKelCard = useMemo(
    () => (!data?.data ? [] : data.data.map(parseDesaKelCard)),
    [data],
  );
  const parsedDesaKelTable = useMemo(
    () => (!data?.data ? [] : data.data.map(parseDesaKelTable)),
    [data],
  );

  const renderData = useMemo(() => {
    const isTableShowData = isShowData && state.dataType === 'TABLE';

    if ((isLoading || isPending) && !isMobile)
      return (
        <SkeletonTable
          columnCount={5}
          rowCount={5}
        />
      );
    if ((isLoading || isPending) && isMobile)
      return <CardListSkeleton count={5} />;

    if ((isLoading || isPending) && state.dataType === 'JSON')
      return <Skeleton className="h-50" />;
    if (isError) return <DesaKelErrorAlert />;

    if (isTableShowData && isMobile)
      return <CardList listData={parsedDesaKelCard} />;
    if (isTableShowData)
      return (
        <Table
          columnHeaders={DesaKelTableHeaders}
          rowData={parsedDesaKelTable}
        />
      );

    if (isShowData && state.dataType === 'JSON')
      return (
        <Code
          content={JSON.stringify(data, null, 2)}
          showCopyButton
        />
      );
    return null;
  }, [
    data,
    isShowData,
    state.dataType,
    isLoading,
    isPending,
    isError,
    isMobile,
    parsedDesaKelTable,
    parsedDesaKelCard,
  ]);

  return <div>{renderData}</div>;
};
