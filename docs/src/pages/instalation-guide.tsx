import Code from '@/components/ui/code';
import { Terminal } from 'lucide-react';

const cloneCommand =
  'git clone --filter=blob:none --no-checkout https://github.com/Yuefii/NusantaraKita.git && cd NusantaraKita && git sparse-checkout set api data && git checkout';

export const InstalationGuide = () => {
  return (
    <article className="text-foreground w-full space-y-10 transition-all">
      {/* Header */}
      <div className="space-y-3">
        <div className="inline-flex items-center gap-2 rounded-md border border-purple-500/20 bg-purple-500/10 px-2.5 py-1 text-xs font-semibold text-purple-600 dark:text-purple-400">
          <Terminal className="h-3.5 w-3.5" />
          <span>Self-Hosting & Local Development</span>
        </div>
        <h1 className="text-foreground text-3xl font-bold tracking-tight sm:text-4xl">
          Panduan Instalasi
        </h1>
        <p className="text-muted-foreground max-w-3xl text-base leading-relaxed sm:text-lg">
          Pelajari cara menjalankan API Nusantara Kita secara mandiri
          (self-host) di lingkungan lokal menggunakan Go runtime atau container
          Docker.
        </p>
      </div>

      <hr className="border-border/60" />

      {/* Step 1: Clone */}
      <section
        id="instalasi-cli"
        className="scroll-mt-20 space-y-4"
      >
        <h2 className="group border-border/50 text-foreground flex items-center gap-2.5 border-b pb-2 text-xl font-bold tracking-tight sm:text-2xl">
          <span className="bg-primary/10 text-primary flex h-6 w-6 items-center justify-center rounded-full text-xs font-bold">
            1
          </span>
          <span>Clone Repository (Sparse Checkout)</span>
          <a
            href="#instalasi-cli"
            className="text-muted-foreground/40 opacity-0 transition-opacity group-hover:opacity-100"
          >
            #
          </a>
        </h2>
        <p className="text-muted-foreground text-sm leading-relaxed sm:text-base">
          Clone repositori Nusantara Kita dengan metode <em>sparse-checkout</em>{' '}
          agar hanya mengunduh direktori <code>api</code> dan <code>data</code>:
        </p>
        <Code
          content={cloneCommand}
          showCopyButton
          language="bash"
        />
      </section>

      {/* Step 2: Environment */}
      <section
        id="konfigurasi-env"
        className="scroll-mt-20 space-y-4"
      >
        <h2 className="group border-border/50 text-foreground flex items-center gap-2.5 border-b pb-2 text-xl font-bold tracking-tight sm:text-2xl">
          <span className="bg-primary/10 text-primary flex h-6 w-6 items-center justify-center rounded-full text-xs font-bold">
            2
          </span>
          <span>Setup Environment</span>
          <a
            href="#konfigurasi-env"
            className="text-muted-foreground/40 opacity-0 transition-opacity group-hover:opacity-100"
          >
            #
          </a>
        </h2>
        <p className="text-muted-foreground text-sm leading-relaxed sm:text-base">
          Masuk ke folder API dan buat salinan berkas konfigurasi lingkungan
          dari template:
        </p>
        <Code
          content={`cd api\ncp .env.example .env`}
          showCopyButton
          language="bash"
        />
      </section>

      {/* Step 3: Run Server */}
      <section
        id="menjalankan-server"
        className="scroll-mt-20 space-y-8"
      >
        <h2 className="group border-border/50 text-foreground flex items-center gap-2.5 border-b pb-2 text-xl font-bold tracking-tight sm:text-2xl">
          <span className="bg-primary/10 text-primary flex h-6 w-6 items-center justify-center rounded-full text-xs font-bold">
            3
          </span>
          <span>Menjalankan Server</span>
          <a
            href="#menjalankan-server"
            className="text-muted-foreground/40 opacity-0 transition-opacity group-hover:opacity-100"
          >
            #
          </a>
        </h2>

        {/* Go Native */}
        <div className="space-y-4">
          <h3 className="text-foreground text-base font-semibold">
            Opsi A: Menjalankan dengan Go Native
          </h3>

          <div className="space-y-4">
            <div>
              <p className="text-muted-foreground mb-1.5 text-xs font-semibold tracking-wider uppercase">
                1. Unduh Dependensi Go Modules
              </p>
              <Code
                content={`go mod download`}
                showCopyButton
                language="bash"
              />
            </div>

            <div>
              <p className="text-muted-foreground mb-1.5 text-xs font-semibold tracking-wider uppercase">
                2. Jalankan HTTP Server
              </p>
              <Code
                content={`go run cmd/server/main.go`}
                showCopyButton
                language="bash"
              />
            </div>

            <div>
              <p className="text-muted-foreground mb-1.5 text-xs font-semibold tracking-wider uppercase">
                3. Build Binary (Opsional untuk Produksi)
              </p>
              <Code
                content={`go build -o server ./cmd/server`}
                showCopyButton
                language="bash"
              />
            </div>

            <div>
              <p className="text-muted-foreground mb-1.5 text-xs font-semibold tracking-wider uppercase">
                4. Jalankan Unit & Integration Test
              </p>
              <Code
                content={`go test ./...`}
                showCopyButton
                language="bash"
              />
            </div>
          </div>
        </div>

        {/* Docker Compose */}
        <div className="space-y-4 pt-2">
          <h3 className="text-foreground text-base font-semibold">
            Opsi B: Menjalankan dengan Docker Compose
          </h3>
          <p className="text-muted-foreground text-sm leading-relaxed sm:text-base">
            Alternatif praktis tanpa perlu menginstal Go runtime di komputer
            lokal:
          </p>

          <Code
            content={`docker-compose up -d`}
            showCopyButton
            language="bash"
          />

          <p className="text-muted-foreground pt-1 text-xs">
            Gunakan flag <code>--build</code> jika ada pembaruan kode:
          </p>
          <Code
            content={`docker-compose up -d --build`}
            showCopyButton
            language="bash"
          />
        </div>
      </section>
    </article>
  );
};
