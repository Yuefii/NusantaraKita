import { useCallback, useMemo } from 'react';
import { useSearchParams } from 'react-router';

const useParams = () => {
  const [searchParams, setSearchParams] = useSearchParams();

  const params = useMemo(() => {
    const pagination = searchParams.get('pagination');
    const limit = searchParams.get('limit');
    const halaman = searchParams.get('halaman');

    return {
      pagination: pagination !== null ? pagination === 'true' : true,
      limit: limit !== null ? parseInt(limit, 10) : 10,
      halaman: halaman !== null ? parseInt(halaman, 10) : 1,
    };
  }, [searchParams]);

  const setParams = useCallback(
    (newParams: Partial<ParamsApi>) => {
      const updatedParams = new URLSearchParams(searchParams);

      if (newParams.limit !== undefined) {
        updatedParams.set('limit', String(newParams.limit));
        updatedParams.set('halaman', '1');
      } else if (newParams.halaman !== undefined) {
        updatedParams.set('halaman', String(newParams.halaman));
      }

      setSearchParams(updatedParams);
    },
    [searchParams, setSearchParams],
  );

  return {
    params,
    setParams,
  };
};

export default useParams;
