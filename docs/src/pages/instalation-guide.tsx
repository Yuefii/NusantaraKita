import Code from '@/components/ui/code';
import { useIsMobile } from '@/hooks/use-mobile';
import { useIsTablet } from '@/hooks/use-tablet';
import { cn } from '@/lib/utils';

const cloneCommand = [
  'git clone --filter=blob:none --no-checkout https://github.com/Yuefii/NusantaraKita.git && cd NusantaraKita && git sparse-checkout set api data && git checkout',
].join('\n');

export const InstalationGuide = () => {
  const isMobile = useIsMobile();
  const isTablet = useIsTablet();
  const isSmallScreen = isMobile || isTablet;

  return (
    <section
      className={cn('max-w-5xl bg-white text-gray-800', {
        'p-5': isSmallScreen,
        'p-8': !isSmallScreen,
      })}
    >
      <h1 className="mb-10 text-3xl font-bold text-gray-700">
        API Documentation - Installation Guide
      </h1>

      {/* Section Clone */}
      <div className="mb-10">
        <h2 className="mb-4 text-xl font-semibold text-gray-600">
          Clone Repository
        </h2>
        <p className="mb-4 leading-relaxed text-gray-600">
          Clone project NusantaraKita ke local kamu menggunakan perintah
          berikut:
        </p>
        <Code
          content={cloneCommand}
          showCopyButton
        />
      </div>

      {/* Section Setup Env */}
      <div className="mb-10 space-y-4">
        <h2 className="text-xl font-semibold text-gray-600">
          Setup Environment
        </h2>
        <Code
          content={`cd api`}
          showCopyButton
        />
        <Code
          content={`cp .env.example .env`}
          showCopyButton
        />
      </div>

      {/* Section Running API with Go */}
      <div className="mb-10">
        <h2 className="mb-4 text-xl font-semibold text-gray-600">
          Menjalankan Project dengan Go
        </h2>

        <div className="mb-6">
          <p className="mb-4 leading-relaxed text-gray-600">
            Unduh dependencies
          </p>
          <Code
            content={`go mod download`}
            showCopyButton
          />
        </div>

        <div className="mb-6">
          <p className="mb-4 leading-relaxed text-gray-600">Jalankan server</p>
          <Code
            content={`go run cmd/server/main.go`}
            showCopyButton
          />
        </div>

        <div className="mb-6">
          <p className="mb-4 leading-relaxed text-gray-600">
            Build project (opsional)
          </p>
          <Code
            content={`go build -o server ./cmd/server`}
            showCopyButton
          />
        </div>

        <div className="mb-6">
          <p className="mb-4 leading-relaxed text-gray-600">Jalankan test</p>
          <Code
            content={`go test ./...`}
            showCopyButton
          />
        </div>
      </div>

      {/* Section Running API with Docker */}
      <div className="mb-10">
        <h2 className="mb-4 text-xl font-semibold text-gray-600">
          Menjalankan Project dengan Docker
        </h2>
        <p className="mb-4 leading-relaxed text-gray-600">
          Setelah clone project kamu bisa menjalankan project dengan mudah
          menggunakan perintah:
        </p>
        <Code
          content={`docker-compose up -d`}
          showCopyButton
        />
        <p className="my-4 leading-relaxed text-gray-600">
          atau gunakan perintah ini jika ada perubahan kode dan kamu ingin
          melakukan <em>build</em> ulang image API-nya:
        </p>
        <Code
          content={`docker-compose up -d --build`}
          showCopyButton
        />
      </div>
    </section>
  );
};
