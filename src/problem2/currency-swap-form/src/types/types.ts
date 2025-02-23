export interface Token {
  currency: string,
  date: string,
  price: number,
}

export type ExchangeRates = {
  [key: string]: number;
};