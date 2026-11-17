import {
  ProfileCard,
  type ProfileCardProps,
} from '@/components/ui/profile-card';
import { BookOpen, ExternalLink } from 'lucide-react';

const developers: ProfileCardProps[] = [
  {
    name: 'M Mupid Ahmadiawan',
    role: 'Backend Developer',
    avatarUrl: '/mupidahmadiawan.jpeg',
    socialLinks: [
      { type: 'github', url: 'https://github.com/Yuefii' },
      {
        type: 'linkedin',
        url: 'https://www.linkedin.com/in/muhamad-mupid-ahmadiawan-3b2a95292/',
      },
    ],
  },
  {
    name: 'Abdillah Juniansyah',
    role: 'Frontend Developer',
    avatarUrl: '/abdillahjuniansyah.jpeg',
    socialLinks: [
      { type: 'github', url: 'https://github.com/dillahCodes' },
      {
        type: 'linkedin',
        url: 'https://www.linkedin.com/in/abdillahjuniansyah/',
      },
    ],
  },
  {
    name: 'Arif Rhamadan',
    role: 'Data Engineer',
    avatarUrl: '/arifrhamadan.jpeg',
    socialLinks: [
      { type: 'github', url: 'https://github.com/arifrhamadan' },
      { type: 'linkedin', url: 'https://www.linkedin.com/in/arifrhamadan' },
    ],
  },
];

export const AboutApi = () => {
  return (
    <article className="text-foreground w-full space-y-10 transition-all">
      {/* Header */}
      <div className="space-y-3">
        <div className="border-primary/25 bg-primary/10 text-primary inline-flex items-center gap-1.5 rounded-md border px-2.5 py-0.5 text-xs font-semibold">
          <BookOpen className="h-3.5 w-3.5" />
          <span>Tentang Inisiatif</span>
        </div>
        <h1 className="text-foreground text-3xl font-bold tracking-tight sm:text-4xl">
          Tentang Nusantara Kita
        </h1>
        <p className="text-muted-foreground text-base leading-relaxed sm:text-lg">
          Nusantara Kita adalah API terbuka yang menyediakan data wilayah
          administratif dan geospasial Indonesia. Dirancang agar mudah
          diintegrasikan ke berbagai jenis aplikasi digital, mulai dari
          e-commerce, logistik, pemerintahan, hingga riset akademik.
        </p>
      </div>

      <hr className="border-border/60" />

      {/* 1. Latar Belakang & Regulasi */}
      <section
        id="latar-belakang"
        className="scroll-mt-20 space-y-4"
      >
        <h2 className="group border-border/50 text-foreground flex items-center gap-2 border-b pb-2 text-xl font-bold tracking-tight sm:text-2xl">
          <span>Sumber Data & Regulasi Resmi</span>
          <a
            href="#latar-belakang"
            className="text-muted-foreground/40 opacity-0 transition-opacity group-hover:opacity-100"
          >
            #
          </a>
        </h2>
        <p className="text-muted-foreground text-sm leading-relaxed sm:text-base">
          Data wilayah administratif pemerintahan Indonesia yang disajikan dalam
          API ini disusun mengacu pada regulasi resmi{' '}
          <span className="text-foreground decoration-primary/40 font-semibold underline underline-offset-4">
            Kepmendagri No 300.2.2-2138 Tahun 2025
          </span>
          . Pembaruan dilakukan berkala demi menjaga validitas dan relevansi
          informasi batas wilayah, koordinat spasial, dan kodepos.
        </p>
      </section>

      {/* 2. Dataset References */}
      <section
        id="cakupan-data"
        className="scroll-mt-20 space-y-4"
      >
        <h2 className="group border-border/50 text-foreground flex items-center gap-2 border-b pb-2 text-xl font-bold tracking-tight sm:text-2xl">
          <span>Dataset & Open Source Attribution</span>
          <a
            href="#cakupan-data"
            className="text-muted-foreground/40 opacity-0 transition-opacity group-hover:opacity-100"
          >
            #
          </a>
        </h2>
        <p className="text-muted-foreground text-sm leading-relaxed sm:text-base">
          Proyek ini terwujud berkat kontribusi dataset terbuka dari komunitas.
          Kami sangat mengapresiasi repositori berikut:
        </p>

        <div className="grid grid-cols-1 gap-3 sm:grid-cols-2">
          <a
            href="https://github.com/cahyadsn/wilayah_boundaries"
            target="_blank"
            rel="noopener noreferrer"
            className="border-border bg-muted/30 hover:border-primary/40 hover:bg-muted/60 group flex items-center justify-between rounded-lg border p-4 transition-all"
          >
            <div className="min-w-0">
              <div className="text-foreground truncate text-sm font-semibold">
                cahyadsn / wilayah_boundaries
              </div>
              <div className="text-muted-foreground mt-0.5 truncate text-xs">
                Dataset batas poligon GeoJSON wilayah Indonesia
              </div>
            </div>
            <ExternalLink className="text-muted-foreground group-hover:text-primary ml-2 h-4 w-4 shrink-0 transition-colors" />
          </a>

          <a
            href="https://github.com/cahyadsn/wilayah_kodepos"
            target="_blank"
            rel="noopener noreferrer"
            className="border-border bg-muted/30 hover:border-primary/40 hover:bg-muted/60 group flex items-center justify-between rounded-lg border p-4 transition-all"
          >
            <div className="min-w-0">
              <div className="text-foreground truncate text-sm font-semibold">
                cahyadsn / wilayah_kodepos
              </div>
              <div className="text-muted-foreground mt-0.5 truncate text-xs">
                Dataset kodepos resmi seluruh desa dan kelurahan
              </div>
            </div>
            <ExternalLink className="text-muted-foreground group-hover:text-primary ml-2 h-4 w-4 shrink-0 transition-colors" />
          </a>
        </div>
      </section>

      {/* 3. Tim Pengembang */}
      <section
        id="arsitektur"
        className="scroll-mt-20 space-y-5"
      >
        <h2 className="group border-border/50 text-foreground flex items-center gap-2 border-b pb-2 text-xl font-bold tracking-tight sm:text-2xl">
          <span>Tim Pengembang</span>
          <a
            href="#arsitektur"
            className="text-muted-foreground/40 opacity-0 transition-opacity group-hover:opacity-100"
          >
            #
          </a>
        </h2>
        <p className="text-muted-foreground text-sm leading-relaxed sm:text-base">
          API Nusantara Kita dikembangkan dan dikelola bersama sebagai proyek
          sumber terbuka (open-source) demi kemudahan integrasi geospasial bagi
          developer Indonesia.
        </p>

        <div className="grid grid-cols-1 gap-6 sm:grid-cols-2 md:grid-cols-3">
          {developers.map((dev) => (
            <ProfileCard
              key={dev.name}
              {...dev}
            />
          ))}
        </div>
      </section>
    </article>
  );
};

export default AboutApi;
