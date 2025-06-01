export function stop(e: React.SyntheticEvent | Event) {
  e.stopPropagation();
}

export const blockKeyEvent = {
  onKeyDown: stop,
  onKeyUp: stop,
  tabIndex: 0,
};

export function prevent(e: React.SyntheticEvent | Event) {
  e.preventDefault();
}

export const preventFormEvent = (event: React.FormEvent<HTMLFormElement>) => {
  event.stopPropagation();
  event.preventDefault();
};
