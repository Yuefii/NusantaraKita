import { FaGithub, FaLinkedin } from 'react-icons/fa';
import { Avatar, AvatarFallback, AvatarImage } from './avatar';

interface SocialLink {
  type: 'github' | 'linkedin';
  url: string;
}

export interface ProfileCardProps {
  name: string;
  role: string;
  avatarUrl: string;
  socialLinks?: SocialLink[];
}

const SocialIcon = ({ type, url }: SocialLink) => {
  const iconMap = {
    github: <FaGithub className="h-5 w-5" />,
    linkedin: <FaLinkedin className="h-5 w-5" />,
  };

  return (
    <a
      href={url}
      target="_blank"
      rel="noopener noreferrer"
      className="text-muted-foreground hover:bg-muted hover:text-foreground rounded-md p-2 transition-all"
      aria-label={type}
    >
      {iconMap[type]}
    </a>
  );
};

export const ProfileCard = ({
  name,
  role,
  avatarUrl,
  socialLinks,
}: ProfileCardProps) => {
  return (
    <div className="group border-border/70 bg-card hover:border-primary/40 relative flex flex-col items-center rounded-lg border p-6 text-center shadow-xs transition-all duration-300 hover:shadow-md">
      <div className="relative mb-4">
        <Avatar className="ring-border/50 group-hover:ring-primary/40 size-24 object-cover ring-4 transition-all">
          <AvatarImage
            src={avatarUrl}
            alt={name}
          />
          <AvatarFallback className="text-lg font-semibold">
            {name.charAt(0)}
          </AvatarFallback>
        </Avatar>
      </div>
      <h4 className="text-foreground text-lg font-bold tracking-tight">
        {name}
      </h4>
      <p className="text-muted-foreground mt-1 text-xs font-medium sm:text-sm">
        {role}
      </p>
      {socialLinks && socialLinks.length > 0 && (
        <div className="border-border/60 mt-4 flex items-center gap-1 border-t pt-4">
          {socialLinks.map((link) => (
            <SocialIcon
              key={link.type}
              {...link}
            />
          ))}
        </div>
      )}
    </div>
  );
};
