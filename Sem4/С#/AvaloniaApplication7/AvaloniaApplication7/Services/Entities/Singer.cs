using SQLite;

namespace AvaloniaApplication7.Services.Entities;



[System.ComponentModel.DataAnnotations.Schema.Table("Singers")]
public class Singer
{
    [PrimaryKey, AutoIncrement, Indexed]
    public int Id { get; set; }
    public string Name { get; set; }
    public string Country { get; set; }
}