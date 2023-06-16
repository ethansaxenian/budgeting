export const strToDate = (dateStr: string): Date => {
  const parts = dateStr.split('-');

  return new Date(
    parseInt(parts[0]),
    parseInt(parts[1]) - 1,
    parseInt(parts[2])
  );
};

export const dateToStr = (date: Date): string => {
  const month = `${date.getMonth() + 1}`.padStart(2, '0');
  const day = `${date.getDate()}`.padStart(2, '0');

  return `${date.getFullYear()}-${month}-${day}`;
};

export const formatAmount = (amount: number): string => {
  if (amount >= 0) {
    return `$${amount}`;
  }

  return `-$${Math.abs(amount)}`;
};

export const colorAmount = (amount: number): string => {
  return amount === 0 ? 'black' : amount > 0 ? 'green' : 'red';
};

export const round = (num: number): number => {
  return Math.round(num * 100) / 100;
};
