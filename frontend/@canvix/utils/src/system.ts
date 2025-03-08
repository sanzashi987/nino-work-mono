type OS = 'win' | 'mac' | 'linux' | 'unix' | undefined;
export const getOS = (): OS => {
  const platform = navigator.userAgent;
  if (platform.includes('Win')) return 'win';
  if (platform.includes('Mac')) return 'mac';
  // if (platform.includes('x11') || platform.includes('unix')) return 'unix';
  if (platform.includes('Linux')) return 'linux';
  return undefined;
};

export const isMac = getOS() === 'mac';

export const sort = (a: number, b: number) => a - b;
