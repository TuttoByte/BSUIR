namespace Lab2.Collections;

public class Patinent
{
    public int Age { get; set; }
    public string Name { get; set; }
    public string Desie{get;set;}
    
    public Patinent(int age, string name, string desie)
    {
        Age = age;
        Name = name;
        Desie = desie;
    }
}