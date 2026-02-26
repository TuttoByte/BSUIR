// See https://aka.ms/new-console-template for more information

using Lab2.Services;
using Lab2.Collections;
using LoremNET;

namespace Lab2;



 class Program
{
    static async Task Main(string[] args)
    {
        
        Console.WriteLine($"Work started {Thread.CurrentThread.ManagedThreadId}");
        

        var progress = new Progress<string>(ThredEndReport);
        var memStram = new MemoryStream();
        var streamServe = new StreamService<Patinent>();


        Func<Patinent, bool> filter = (Patinent p) =>
        {
            if (p.Desie == "Asthma")
            {
                return true;
            }
            return false;
        };


        var patientList = GenerateRandom();
        
        
        
        
        
       Task task1 = Task.Run(()=>streamServe.WriteToStreamAync(memStram, patientList,  progress));
       Task task2 = task1.ContinueWith((Task t) => streamServe.CopyFromStreamAsync(memStram, "/home/udainoko/Documents/BSUIR/Sem4/С#/Lab2/Lab2/FIles/Test.txt", progress));

       Task.WaitAll(task2);
       

        int num = await streamServe.GetStatisticsAsync("/home/udainoko/Documents/BSUIR/Sem4/С#/Lab2/Lab2/FIles/Test.txt", filter);
        Console.WriteLine($"Number of patients with Asthma: {num}");
    }
    
    
    private static void ThredEndReport(string action)
    {
        Console.WriteLine(action);
    }

    private static List<Patinent> GenerateRandom()
    {
        string[] names =
        {
            "Adam", "Anny", "Joseph", "Robbin", "Alex"
        };
        string[] desies =
        [
            "Flu", "Cold", "Tonsillitis", "Pneumonia", "Bronchitis", "Gastroenteritis", "Conjunctivitis", "Otitis",
            "Sinusitis", "Tonsillitis", "Dermatitis", "Eczema", "Psoriasis", "Arthritis", "Gastritis", "Stomach Ulcer",
            "Hepatitis", "Diabetes", "Hypertension", "Asthma"
        ];

      
        var persons = new List<Patinent>();

        for (int i = 0; i < 1000; i++)
        {   
            var ages = Convert.ToInt32(LoremNET.Lorem.Number(10, 100));
            persons.Add(new Patinent(ages, LoremNET.Lorem.Random(names), Lorem.Random(desies)));
        }
        return persons;
    }
}


