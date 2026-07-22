import en from './en';
import fr from './fr';
import type { Dict } from './types';

export type Lang = 'en' | 'fr';
export const dictionaries: Record<Lang, Dict> = { en, fr };
export type { Dict };

class LangState {
  current = $state<Lang>('en');
  toggle() {
    this.current = this.current === 'en' ? 'fr' : 'en';
  }
  get t(): Dict {
    return dictionaries[this.current];
  }
}

export const lang = new LangState();
