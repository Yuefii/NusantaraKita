import axiosInstance from '@/lib/axios';
import { keepPreviousData, useQuery } from '@tanstack/react-query';

export const getProvinsi = async (
  params?: ParamsApi,
): Promise<ProvinsiApiRes> => {
  const response = await axiosInstance.get('provinsi', { params });
  return response.data as ProvinsiApiRes;
};

const useGetProvinsi = (
  options: ParamsApi,
  queryOptions?: { enabled?: boolean },
) => {
  const {
    data,
    isLoading,
    isError,
    isSuccess,
    isPending,
    refetch,
    isFetching,
  } = useQuery({
    queryKey: ['provinsi', options],
    queryFn: () => getProvinsi(options),
    placeholderData: keepPreviousData,
    refetchOnWindowFocus: false,
    retry: false,
    enabled: queryOptions?.enabled ?? true,
  });

  return {
    data,
    isLoading,
    isError,
    isSuccess,
    isPending,
    refetch,
    isFetching,
  };
};

export default useGetProvinsi;
