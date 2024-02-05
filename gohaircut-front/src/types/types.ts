export type Client = {
  uid: string;
  email: string;
  age: number;
  password: string;
};

export type HairSaloon = {
  uid: string;
  name: string;
  address: string;
  email: string;
  phone: string;
  openingTime: string;
  closingTime: string;
};

export type Hairdresser = {
  uid: string;
  saloonId: HairSaloon | null;
  firstName: string;
  speciality: string;
};

export type Schedule = {
  uid: string;
  hairdresserId: Hairdresser | null;
  startHour: string;
  endHour: string;
  availability: boolean;
};

export type Appointments = {
  uid: string;
  saloonId: HairSaloon | null;
  clientId: Client | null;
  hairdresserId: Hairdresser | null;
  startHour: string;
  status: string;
};

export type Admin = {
  uid: string;
  name: string;
  email: string;
  password: string;
};
