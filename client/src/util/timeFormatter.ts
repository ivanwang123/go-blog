const months = [
  "January",
  "February",
  "March",
  "April",
  "May",
  "June",
  "July",
  "August",
  "September",
  "October",
  "November",
  "December",
];

export const formatTimestamp = (timestamp: string) => {
  let date = new Date(timestamp);
  return `${months[date.getMonth()]} ${date.getDate()}, ${date.getFullYear()}`;
};
