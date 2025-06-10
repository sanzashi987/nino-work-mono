export const saturate = (lower: number, upper: number, input: number): number =>
  // eslint-disable-next-line no-nested-ternary
  input < lower ? lower : input > upper ? upper : input;
