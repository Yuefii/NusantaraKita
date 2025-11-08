import { Alert, AlertDescription, AlertTitle } from '@/components/ui/alert';
import { CardList } from '@/components/ui/card-list';
import { useResponseType } from '@/components/ui/response-type';
import { CardListSkeleton } from '@/components/ui/skelekton-card-list';
import SkeletonTable from '@/components/ui/skelekton-table';
import { Skeleton } from '@/components/ui/skeleton';
import { useIsMobile } from '@/hooks/use-mobile';
import { Globe } from 'lucide-react';
import { useMemo } from 'react';
import {
  FaArrowsAltH,
  FaArrowsAltV,
  FaHashtag,
  FaMapMarkerAlt,
} from 'react-icons/fa';
import { MdErrorOutline } from 'react-icons/md';
import Code from '../../ui/code';
import Table from '../../ui/table';

interface ProvinsiDataProps {
  data: ProvinsiApiRes | undefined;
  isError: boolean;
  isLoading: boolean;
  isSuccess: boolean;
  isPending: boolean;
}

const provinsiTableHeaders = [
  'Kode',
  'Provinsi',
  'Latitude',
  'Longitude',
  'geojson_url',
];

const parseProvinsi = (provinsi: ProvinsiApi) => ({
  Kode: provinsi.kode,
  Provinsi: provinsi.nama,
  Latitude: provinsi.lat,
  Longitude: provinsi.lng,
  geojson_url: provinsi.geojson_url,
});

const parseProvinsiCard = (provinsi: ProvinsiApi) => [
  {
    icon: FaHashtag,
    title: 'Kode',
    value: provinsi.kode.toString(),
  },
  {
    icon: FaMapMarkerAlt,
    title: 'Provinsi',
    value: provinsi.nama,
  },
  {
    icon: FaArrowsAltV,
    title: 'Latitude',
    value: provinsi.lat.toString(),
  },
  {
    icon: FaArrowsAltH,
    title: 'Longitude',
    value: provinsi.lng.toString(),
  },
  {
    icon: Globe,
    title: 'geojson_url',
    value: provinsi.geojson_url.toString(),
  },
];

const ProvinsiErrorAlert = () => {
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

export const ProvinsiDisplayData = ({
  data,
  isError,
  isLoading,
  isSuccess,
  isPending,
}: ProvinsiDataProps) => {
  const isMobile = useIsMobile();

  const { state } = useResponseType();
  const isShowData = isSuccess && data?.data;

  const parsedProvinsiCard = useMemo(
    () => (!data?.data ? [] : data.data.map(parseProvinsiCard)),
    [data],
  );
  const parsedProvinsiTable = useMemo(
    () => (!data?.data ? [] : data.data.map(parseProvinsi)),
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
    if (isError) return <ProvinsiErrorAlert />;

    if (isTableShowData && isMobile)
      return <CardList listData={parsedProvinsiCard} />;
    if (isTableShowData)
      return (
        <Table
          columnHeaders={provinsiTableHeaders}
          rowData={parsedProvinsiTable}
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
    parsedProvinsiTable,
    parsedProvinsiCard,
  ]);

  return <div>{renderData}</div>;
};
