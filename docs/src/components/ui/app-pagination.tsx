'use client';

import {
  Pagination,
  PaginationContent,
  PaginationEllipsis,
  PaginationItem,
  PaginationLink,
  PaginationNext,
  PaginationPrevious,
} from '@/components/ui/pagination';
import { useIsMobile } from '@/hooks/use-mobile';
import useParams from '@/hooks/use-params';
import { useCallback, useMemo } from 'react';

interface PaginationProps {
  currentPage: number;
  totalPages: number;
}

type PaginationItemType =
  { type: 'page'; value: number } | { type: 'ellipsis'; value: string };

function shouldShowPreviousButton(
  isFirstPage: boolean,
  isMobile: boolean,
): boolean {
  return !isFirstPage && !isMobile;
}

function shouldShowNextButton(isLastPage: boolean, isMobile: boolean): boolean {
  return !isLastPage && !isMobile;
}

function buildPaginationItems(
  currentPage: number,
  totalPageCount: number,
  isMobile: boolean,
): PaginationItemType[] {
  const items: PaginationItemType[] = [];
  const maxVisible = isMobile ? 3 : 5;
  const halfVisible = Math.floor(maxVisible / 2);

  let startPage = Math.max(1, currentPage - halfVisible);
  let endPage = Math.min(totalPageCount, currentPage + halfVisible);

  if (currentPage - halfVisible < 1) {
    endPage = Math.min(totalPageCount, maxVisible);
  }
  if (currentPage + halfVisible > totalPageCount) {
    startPage = Math.max(1, totalPageCount - maxVisible + 1);
  }

  if (startPage > 1) items.push({ type: 'page', value: 1 });
  if (startPage > 2) items.push({ type: 'ellipsis', value: 'start-ellipsis' });

  for (let i = startPage; i <= endPage; i++)
    items.push({ type: 'page', value: i });

  if (endPage < totalPageCount - 1)
    items.push({ type: 'ellipsis', value: 'end-ellipsis' });
  if (endPage < totalPageCount)
    items.push({ type: 'page', value: totalPageCount });

  return items;
}

function renderPaginationItem(
  item: PaginationItemType,
  index: number,
  currentPage: number,
  onPageChange: (page: number) => void,
) {
  const key = `${item.type}-${item.value}-${index}`;

  if (item.type === 'ellipsis') {
    return (
      <PaginationItem key={key}>
        <PaginationEllipsis />
      </PaginationItem>
    );
  }

  const isActive = item.value === currentPage;
  const linkClassName = isActive
    ? 'hover:bg-black/70 bg-black text-white hover:text-white'
    : 'hover:bg-black hover:text-white';

  return (
    <PaginationItem key={key}>
      <PaginationLink
        onClick={() => onPageChange(item.value)}
        className={linkClassName}
        isActive={isActive}
        aria-current={isActive ? 'page' : undefined}
      >
        {item.value}
      </PaginationLink>
    </PaginationItem>
  );
}

export default function AppPagination({
  currentPage,
  totalPages,
}: PaginationProps) {
  const isMobile = useIsMobile();
  const { setParams } = useParams();

  const actualTotalPages = Math.max(1, totalPages);
  const isFirstPage = currentPage === 1;
  const isLastPage = currentPage === actualTotalPages;

  const navigateToPage = useCallback(
    (page: number) => setParams({ halaman: page }),
    [setParams],
  );

  const goToNextPage = useCallback(() => {
    return navigateToPage(Math.min(currentPage + 1, actualTotalPages));
  }, [navigateToPage, currentPage, actualTotalPages]);

  const goToPreviousPage = useCallback(
    () => navigateToPage(Math.max(currentPage - 1, 1)),
    [navigateToPage, currentPage],
  );

  const paginationItemsToRender = useMemo(() => {
    return buildPaginationItems(currentPage, actualTotalPages, isMobile);
  }, [currentPage, actualTotalPages, isMobile]);

  const renderedPaginationLinks = useMemo(() => {
    return paginationItemsToRender.map((item, index) =>
      renderPaginationItem(item, index, currentPage, navigateToPage),
    );
  }, [paginationItemsToRender, currentPage, navigateToPage]);

  if (actualTotalPages <= 1) return null;

  return (
    <Pagination className="select-none">
      <PaginationContent className="gap-1 sm:gap-2">
        {shouldShowPreviousButton(isFirstPage, isMobile) && (
          <PaginationItem>
            <PaginationPrevious
              onClick={goToPreviousPage}
              className="hover:bg-black hover:text-white"
              aria-disabled={isFirstPage}
            />
          </PaginationItem>
        )}

        {renderedPaginationLinks}

        {shouldShowNextButton(isLastPage, isMobile) && (
          <PaginationItem>
            <PaginationNext
              onClick={goToNextPage}
              className="hover:bg-black hover:text-white"
              aria-disabled={isLastPage}
            />
          </PaginationItem>
        )}
      </PaginationContent>
    </Pagination>
  );
}

export function SkeletonPagination() {
  return (
    <Pagination className="animate-pulse select-none">
      <PaginationContent className="gap-1 sm:gap-2">
        {/* Previous Button Skeleton */}
        <PaginationItem>
          <div className="h-9 w-9 rounded-md bg-gray-200" />
        </PaginationItem>

        {/* Page Number Skeletons */}
        {[...Array(2)].map((_, index) => (
          <PaginationItem key={index}>
            <div className="h-9 w-9 rounded-md bg-gray-200" />
          </PaginationItem>
        ))}

        {/* Next Button Skeleton */}
        <PaginationItem>
          <div className="h-9 w-9 rounded-md bg-gray-200" />
        </PaginationItem>
      </PaginationContent>
    </Pagination>
  );
}
