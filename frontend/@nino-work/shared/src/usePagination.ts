import { useState, useEffect, useCallback, useRef, useMemo } from 'react';
import { PageSize } from './types';

type PaginationRequest<F> = PageSize & Partial<F>;

type PaginationResponse<T> = {
  data: T[];
  total: number;
} & PageSize;

type UsePaginationResult<T, F> = {
  loading: boolean
  data: T[];
  total: number;
  currentPage: number;
  pageSize: number;
  setPage: (page: number) => void;
  setPageSize: (size: number) => void;
  setFilters: (filters: Partial<F>) => void;
  refresh: () => void;
};

export const usePagination = <T, F extends object>(
  requestFn: (params: PaginationRequest<F>) => Promise<PaginationResponse<T>>,
  { page, size, ...others }: Partial<PaginationRequest<F>> = {}
): UsePaginationResult<T, F> => {
  const [pagination, setPagination] = useState<PageSize>({
    page: page ?? 1,
    size: size ?? 25
  });
  const [data, setData] = useState<T[]>([]);
  const [total, setTotal] = useState(0);
  const [filters, setFilters] = useState(others as Partial<F>);
  const [loading, setLoading] = useState(false);
  const requestFnRef = useRef(requestFn);

  requestFnRef.current = requestFn;

  const fetchData = useCallback(async () => {
    setLoading(true);
    try {
      const response = await requestFnRef.current({ ...pagination, ...filters });
      setPagination({ page: response.page, size: response.size });
      setData(response.data);
      setTotal(response.total);
    } catch (error) {
      console.error('Error fetching paginated data:', error);
    } finally {
      setLoading(false);
    }
  }, [pagination, filters]);

  const fetchDataRef = useRef(fetchData);

  useEffect(() => {
    fetchData();
  }, [fetchData]);

  const { setPage, setPageSize, refresh } = useMemo(() => ({
    setPage: (nextPage: number) => {
      setPagination((prev) => ({ ...prev, page: nextPage }));
    },
    setPageSize: (nextSize: number) => {
      setPagination((prev) => ({ ...prev, size: nextSize, page: 1 })); // Reset to page 1 when page size changes
    },
    refresh: () => {
      fetchDataRef.current();
    }
  }), []);

  return {
    loading,
    data,
    total,
    currentPage: pagination.page,
    pageSize: pagination.size,
    setPage,
    setPageSize,
    setFilters: (newFilters) => setFilters((prev) => ({ ...prev, ...newFilters })),
    refresh
  };
};
