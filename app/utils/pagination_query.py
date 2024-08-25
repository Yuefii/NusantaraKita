class Pagination:
    def __init__(self, query, page=1, per_page=10):
        self.query = query
        self.page = self._validate_page(page)
        self.per_page = self._validate_per_page(per_page)

    def _validate_page(self, page):
        try:
            page = int(page)
        except ValueError:
            raise ValueError("Page must be an integer")
        if page < 1:
            raise ValueError("Page number must be greater than zero")
        return page

    def _validate_per_page(self, per_page):
        try:
            per_page = int(per_page)
        except ValueError:
            raise ValueError("per_page must be an integer")
        if per_page <= 0:
            raise ValueError("per_page must be greater than zero")
        return per_page

    def get_paginated_data(self):
        total_items = self.query.count()
        total_pages = (total_items + self.per_page - 1) // self.per_page

        if self.page > total_pages:
            raise ValueError("Page number exceeds total pages")

        paginated_query = self.query.paginate(
            page=self.page, per_page=self.per_page, error_out=False
        )

        return {
            "pagination": {
                "total_items": paginated_query.total,
                "total_pages": paginated_query.pages,
                "current_page": paginated_query.page,
                "per_page": paginated_query.per_page,
            },
            "data": [
                {"code": item.code, "name": item.name} for item in paginated_query.items
            ],
        }
