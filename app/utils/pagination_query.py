class Pagination:
    """
    Class untuk menangani pagination hasil dari query.

    Attributes:
        query: object yang akan dipagination.
        page: Number halaman yang saat ini digunakan (index dimulai dari 1).
        per_page: jumlah item per halaman.

    Methods:
        __init__(query, page=1, per_page=10):
            Menginisialisasi objek Pagination dengan query, number halaman, dan nilai per_page yang diberikan.
        _validate_page(page):
            Memvalidasi dan mengembalikan number halaman.
        _validate_per_page(per_page):
            Memvalidasi dan mengembalikan jumlah item per halaman.
        get_paginated_data():
            Mengembalikan dictionary yang berisi detail pagination dan data halaman saat ini.
    """

    def __init__(self, query, page=1, per_page=10):
        """
        Menginisialisasi objek Pagination dengan query, number halaman, dan jumlah item perhalaman yang diberikan.

        Args:
            query: Objek query yang akan dipagination. Objek ini harus memiliki method `count()` dan `paginate()`.
            page (int): Number halaman saat ini. Defaultnya adalah 1.
            per_page (int): Jumlah item per halaman. Defaultnya adalah 10.

        Raises:
            ValueError: Jika page atau per_page bukan int positif.
        """
        self.query = query
        self.page = self._validate_page(page)
        self.per_page = self._validate_per_page(per_page)

    def _validate_page(self, page):
        """
        Memvalidasi nomor halaman.

        Args:
            page: Number halaman yang akan divalidasi.

        Returns:
            int: Nomor halaman yang telah divalidasi.

        Raises:
            ValueError: Jika page bukan int atau kurang dari 1.
        """
        try:
            page = int(page)
        except ValueError:
            raise ValueError("Page must be an integer")
        if page < 1:
            raise ValueError("Page number must be greater than zero")
        return page

    def _validate_per_page(self, per_page):
        """
        Memvalidasi jumlah item per halaman.

        Args:
            per_page: Jumlah item per halaman yang akan divalidasi.

        Returns:
            int: Jumlah item per halaman yang sudah divalidasi.

        Raises:
            ValueError: Jika per_page bukan int atau kurang dari, sama dengan 0.
        """
        try:
            per_page = int(per_page)
        except ValueError:
            raise ValueError("per_page must be an integer")
        if per_page <= 0:
            raise ValueError("per_page must be greater than zero")
        return per_page

    def get_paginated_data(self):
        """
        Mengambil data yang dipagination berdasarkan halaman saat ini dan jumlah item per halaman.

        Returns:
            dict: Dictionary yang berisi informasi pagination dan data halaman saat ini.
                Dictionary ini mencakup:
                - "pagination": Dictionary dengan jumlah total item, total jumlah halaman, number halaman saat ini, dan item per halaman.
                - "data": Daftar dictionary yang mewakili item yang dipagination, masing-masing memiliki attribute "code" dan "name".

        Raises:
            ValueError: Jika number halaman saat ini melebihi total jumlah halaman.
        """
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
