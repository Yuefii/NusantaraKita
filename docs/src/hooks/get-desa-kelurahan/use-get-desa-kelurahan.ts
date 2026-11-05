import axiosInstace from '@/lib/axios';
import { keepPreviousData, useQuery } from '@tanstack/react-query';

export const getDesaKelurahan = async (
  params?: ParamsApi,
): Promise<DesaKelApiRes> => {
  const response = await axiosInstace.get('desa-kel', { params });
  return response.data as DesaKelApiRes;
};

const useGetDesaKelurahan = (options: ParamsApi) => {
  const { data, isLoading, isError, isSuccess, isPending } = useQuery({
    queryKey: ['desa-kel', options],
    queryFn: () => getDesaKelurahan(options),
    placeholderData: keepPreviousData,
    refetchOnWindowFocus: false,
    retry: false,
  });

  return { data, isLoading, isError, isSuccess, isPending };
};

export default useGetDesaKelurahan;
