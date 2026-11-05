import axiosInstace from '@/lib/axios';
import { keepPreviousData, useQuery } from '@tanstack/react-query';

export const getKecamatan = async (
  params?: ParamsApi,
): Promise<KecamatanApiRes> => {
  const response = await axiosInstace.get('kecamatan', { params });
  return response.data as KecamatanApiRes;
};

const useGetKecamatan = (options: ParamsApi) => {
  const { data, isLoading, isError, isSuccess, isPending } = useQuery({
    queryKey: ['kecamatan', options],
    queryFn: () => getKecamatan(options),
    placeholderData: keepPreviousData,
    refetchOnWindowFocus: false,
    retry: false,
  });

  return { data, isLoading, isError, isSuccess, isPending };
};

export default useGetKecamatan;
