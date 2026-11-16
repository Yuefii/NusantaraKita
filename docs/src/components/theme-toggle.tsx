import { useTheme } from '@/context/theme-provider';
import { Monitor, Moon, Sun } from 'lucide-react';
import { Button } from './ui/button';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from './ui/dropdown-menu';

export function ThemeToggle() {
  const { theme, setTheme, resolvedTheme } = useTheme();

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button
          variant="outline"
          size="icon"
          className="border-border/60 bg-background/50 hover:bg-muted text-foreground relative h-9 w-9 cursor-pointer rounded-md transition-all"
          aria-label="Toggle theme"
        >
          {resolvedTheme === 'dark' ? (
            <Moon className="h-4 w-4 transition-transform duration-200" />
          ) : (
            <Sun className="h-4 w-4 transition-transform duration-200" />
          )}
          <span className="sr-only">Toggle theme</span>
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent
        align="end"
        className="bg-popover/95 border-border/80 w-36 backdrop-blur-md"
      >
        <DropdownMenuItem
          onClick={() => setTheme('light')}
          className="flex cursor-pointer items-center gap-2"
        >
          <Sun className="h-4 w-4 text-amber-500" />
          <span
            className={theme === 'light' ? 'text-primary font-semibold' : ''}
          >
            Light
          </span>
        </DropdownMenuItem>
        <DropdownMenuItem
          onClick={() => setTheme('dark')}
          className="flex cursor-pointer items-center gap-2"
        >
          <Moon className="h-4 w-4 text-sky-400" />
          <span
            className={theme === 'dark' ? 'text-primary font-semibold' : ''}
          >
            Dark
          </span>
        </DropdownMenuItem>
        <DropdownMenuItem
          onClick={() => setTheme('system')}
          className="flex cursor-pointer items-center gap-2"
        >
          <Monitor className="text-muted-foreground h-4 w-4" />
          <span
            className={theme === 'system' ? 'text-primary font-semibold' : ''}
          >
            System
          </span>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
